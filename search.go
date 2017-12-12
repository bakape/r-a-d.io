package main

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/bakape/r-a-d.io/templates"
	"gopkg.in/olivere/elastic.v5"
)

var (
	elasticClient *elastic.Client
)

// Connect to Elastic Search
func initElastic() (err error) {
	elasticClient, err = elastic.NewClient()
	return
}

// Query Elastic Search for matching songs
func querySearch(query string, page int, ctx context.Context) (
	buf []byte, err error,
) {
	// Testing without ES running
	if elasticClient == nil {
		return []byte("TEST: no search possible"), nil
	}

	q := elastic.NewQueryStringQuery(query).
		DefaultOperator("AND")
	for _, s := range [...]string{"title", "artist", "album", "tags", "_id"} {
		q.Field(s)
	}
	res, err := elasticClient.Search().
		Index("song-database").
		Type("track").
		Query(q).
		SortBy(
			elastic.NewFieldSort("priority").Desc(),
			elastic.NewFieldSort("_score").Desc(),
		).
		PostFilter(
			elastic.NewBoolQuery().
				Must(elastic.NewTermQuery("usable", 1)),
		).
		From(page * 20).
		Size(20).
		Do(ctx)
	if err != nil {
		return
	}
	buf, err = json.Marshal(res)
	return
}

// Serve search result page
func serveSearch(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	page, _ := strconv.ParseUint(q.Get("page"), 10, 64)
	buf, err := querySearch(q.Get("q"), int(page), r.Context())
	if err != nil {
		text500(w, r, err)
		return
	}

	w.Write([]byte(templates.Search(buf)))
}
