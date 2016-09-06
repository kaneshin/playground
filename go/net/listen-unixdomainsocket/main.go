package main

import (
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func finalize(ln net.Listener) {
	if err := ln.Close(); err != nil {
		log.Fatalf("error %v", err)
	}
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("It works!"))
	})

	ln, err := net.Listen("unix", "/var/run/gopher/go.sock")
	if err != nil {
		log.Fatalf("error %v", err)
	}
	defer func(ln net.Listener) {
		if err := ln.Close(); err != nil {
			log.Fatalf("error %v", err)
		}
	}(ln)

	c := make(chan os.Signal, 2)
	go func(ln net.Listener, c chan os.Signal) {
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		<-c
		if err := ln.Close(); err != nil {
			log.Fatalf("error %v", err)
		}
		os.Exit(1)
	}(ln, c)

	log.Fatalf("error %v", http.Serve(ln, mux))
}
