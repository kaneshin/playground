package main

import (
	"context"
	"net/http"
	"time"
)

func concat(ctx context.Context, a, b string) string {
	time.Sleep(2 * time.Second)
	time
	return a + " " + b
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ctx1, cancel := context.WithCancel(r.Context())

		var str1, str2, str3 string

		ctx2, _ := context.WithCancel(ctx1)
		go func(ctx context.Context) {
			str1 = concat(ctx, "Hello", "world")
		}(ctx2)

		go func(ctx context.Context) {
			str2 = concat(ctx, "こんにちは", "世界")
		}(ctx2)

		select {
		case <-ctx2.Done():
			str3 = concat(ctx1, str1, str2)
		}

		w.Write([]byte(str3))
		cancel()
	})

	http.ListenAndServe(":9090", nil)
}
