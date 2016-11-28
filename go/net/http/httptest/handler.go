package main

import (
	"log"
	"net/http"
)

func pingHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
	}
}

func echoHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(r.URL.Query().Get("msg")))
	}
}

func init() {
	http.HandleFunc("/ping", pingHandler())
	http.HandleFunc("/echo", echoHandler())
}

func main() {
	log.Fatal(http.ListenAndServe(":8080", nil))
}
