// +build do3

package main

import (
	"log"

	"github.com/kaneshin/playground/go/concurrency-pattern/payload"
)

var queue chan payload.Instance

func start() {
	for {
		select {
		case instance := <-queue:
			body, err := instance.Get()
			if err != nil {
				log.Println(err)
			} else {
				log.Printf("Get %s  (%d bytes)\n", instance.URL, len(body))
			}
		}
	}
}

func init() {
	queue = make(chan payload.Instance, numMaxQueues())
	go start()
}

func do(instance payload.Instance) {
	queue <- instance
}
