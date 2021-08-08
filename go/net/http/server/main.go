package main

import (
	"net/http"
)

const csv = `id,name
1,foo
2,bar
3,baz
`

func main() {
	http.HandleFunc("/text-plain/sample", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("content-type", "text/plain")
		w.Write([]byte(csv))
	})

	http.HandleFunc("/text-html/sample", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("content-type", "text/html")
		w.Write([]byte(csv))
	})

	http.HandleFunc("/text-csv/sample", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("content-type", "text/csv")
		w.Write([]byte(csv))
	})

	http.HandleFunc("/download/sample", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("content-type", "text/csv")
		w.Header().Add("content-disposition", "inline; filename=\"foobar.csv\"")
		w.Write([]byte(csv))
	})

	http.HandleFunc("/octet-stream/sample", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("content-type", "application/octet-stream")
		w.Write([]byte(csv))
	})

	http.ListenAndServe(":8080", nil)
}
