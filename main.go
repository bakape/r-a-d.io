package main

import (
	"flag"
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

	r := httptreemux.NewContextMux()
	r.NotFoundHandler = func(w http.ResponseWriter, _ *http.Request) {
		text404(w)
	}
	r.PanicHandler = text500

	r.GET("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(templates.Index()))
	})
	r.GET("/ass/*path", func(w http.ResponseWriter, r *http.Request) {
		path := filepath.Join("www", extractParam(r, "path"))
		http.ServeFile(w, r, filepath.Clean(path))
	})

	log.Printf("listening on %s\n", *addr)
	if err := http.ListenAndServe(*addr, r); err != nil {
		log.Fatal(err)
	}
}
