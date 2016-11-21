package hello

import (
	"net/http"
)

func init() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			w.Write([]byte("Hello, world!"))

		case http.MethodPost:
			r.ParseForm()
			w.Write([]byte("Hello, Text! " + r.PostForm.Get("text")))

		}
	})
}
