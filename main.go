package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/dimfeld/httptreemux"
)

func main() {
	addr := flag.String("a", ":8010", "server listening address")
	flag.Parse()

	r := httptreemux.NewContextMux()
	r.NotFoundHandler = func(w http.ResponseWriter, _ *http.Request) {
		text404(w)
	}
	r.PanicHandler = text500

	r.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello!")
	})

	log.Printf("listening on %s\n", *addr)
	if err := http.ListenAndServe(*addr, r); err != nil {
		log.Fatal(err)
	}
}
