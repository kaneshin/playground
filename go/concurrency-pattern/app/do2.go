// +build do2

package main

import (
	"log"

	"github.com/kaneshin/playground/go/concurrency-pattern/payload"
)

func do(instance payload.Instance) {
	go func() {
		body, err := instance.Get()
		if err != nil {
			log.Println(err)
		} else {
			log.Printf("Get %s  (%d bytes)\n", instance.URL, len(body))
		}
	}()
}
