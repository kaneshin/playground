package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const (
	sock = "/var/run/gopher/go.sock"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>It works!</h1>\n")
	})

	mux.HandleFunc("/header", func(w http.ResponseWriter, r *http.Request) {
		b, err := json.MarshalIndent(r.Header, "", "  ")
		if err != nil {
			fmt.Fprintf(w, err.Error())
		} else {
			w.Header().Add("Content-Type", "application/json")
			fmt.Fprintf(w, "%v\n", string(b))
		}
	})

	listener, err := net.Listen("unix", sock)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	defer func() {
		if err := listener.Close(); err != nil {
			log.Println("Error:", err.Error())
		}
	}()
	shutdown(listener)
	if err := http.Serve(listener, mux); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func shutdown(listener net.Listener) {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		s := <-c
		log.Println("Got signal:", s)
		if err := listener.Close(); err != nil {
			log.Println("Error:", err.Error())
		}
		os.Exit(1)
	}()
}
