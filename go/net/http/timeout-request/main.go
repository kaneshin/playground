package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"time"
)

var (
	timeout = flag.Float64("timeout", 1.0, "")
	load    = flag.Float64("load", 0.5, "")
)

func main() {
	flag.Parse()

	// start server by httptest
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("It works!"))

		// simulate heavy request
		time.Sleep(time.Duration(*load*1000) * time.Millisecond)
	}))

	// create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(*timeout*1000)*time.Millisecond)
	// gracefully cancel
	defer cancel()

	req, err := http.NewRequest("GET", s.URL, nil)
	if err != nil {
		log.Fatalf("error %v", err)
	}
	req.Cancel = ctx.Done()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("error %v", err)
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(b))
}
