package main

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/bakape/r-a-d.io/common"
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
	songs []common.SearchSong, pages int, err error,
) {
	// Testing without ES running
	if elasticClient == nil {
		return
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

	pages = int(res.Hits.TotalHits) / 20
	if pages == 0 {
		pages = 1
	}

	if res.Hits.TotalHits > 0 {
		songs = make([]common.SearchSong, 0, len(res.Hits.Hits))
		for _, hit := range res.Hits.Hits {
			var song common.SearchSong
			err = json.Unmarshal(*hit.Source, &song)
			if err != nil {
				return
			}
			songs = append(songs, song)
		}
	}

	return
}

// Serve search result page
func serveSearch(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	page, _ := strconv.ParseUint(q.Get("page"), 10, 64)
	query := q.Get("q")
	songs, pages, err := querySearch(query, int(page), r.Context())
	if err != nil {
		text500(w, r, err)
		return
	}

	w.Write([]byte(templates.Search(query, int(page), pages, songs)))
}
