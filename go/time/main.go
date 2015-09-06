package main

import (
	"fmt"
	"time"
)

func main() {
	t, err := time.Parse("2006-01-02T15:04:05Z", "2015-04-01T19:35:49Z")
	if err != nil {
		panic(err)
	}
	fmt.Println(t)
	fmt.Println(t.UnixNano())
	fmt.Println(int64(time.Hour))
}
