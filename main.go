package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/bakape/r-a-d.io/templates"
	"github.com/dimfeld/httptreemux"
)

// TODO: Config reloading. Probably from JSON config file.

func main() {
	addr := flag.String("a", ":8010", "server listening address")
	flag.Parse()

	if err := initElastic(); err != nil {
		// So you can still test some things without elastic running
		log.Println(err)
	}

	r := httptreemux.NewContextMux()
	r.NotFoundHandler = func(w http.ResponseWriter, _ *http.Request) {
		text404(w)
	}
	r.PanicHandler = text500

	r.GET("/", serveIndex)
	r.GET("/search", serveSearch)
	r.GET("/ass/*path", func(w http.ResponseWriter, r *http.Request) {
		path := filepath.Join("www", extractParam(r, "path"))
		http.ServeFile(w, r, filepath.Clean(path))
	})

	log.Printf("listening on %s\n", *addr)
	if err := http.ListenAndServe(*addr, r); err != nil {
		log.Fatal(err)
	}
}

func serveIndex(w http.ResponseWriter, r *http.Request) {
	apiMu.RLock()
	data := apiData
	hash := apiHash
	apiMu.RUnlock()

	etag := fmt.Sprintf(`W/"%s"`, hash)
	if etag == r.Header.Get("If-None-Match") {
		w.WriteHeader(304)
		return
	}
	w.Header().Set("ETag", etag)

	w.Write([]byte(templates.Index(data)))
}
