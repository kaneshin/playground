package main

import (
	"log"
	"time"
)

func init() {
	log.SetFlags(log.LstdFlags)
	log.SetPrefix("[INFO]")
}

func main() {
	for i := 1; i <= 100; i++ {
		log.Printf("%v", map[string]interface{}{
			"host":   "127.0.0.1",
			"time":   time.Now(),
			"status": 200,
		})
	}
}
