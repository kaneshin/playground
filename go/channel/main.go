package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var (
	port = flag.String("port", "8080", "port")
)

func main() {
	hash := map[string]interface{}{}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			id := r.FormValue("id")
			w.Write([]byte(id))
		case "POST":
		case "PUT":
		}
	})

	if *port == "" {
		*port = "8080"
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", *port), nil))
}
