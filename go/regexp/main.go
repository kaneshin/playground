package main

import (
	"fmt"
	"regexp"
	"sync"
)

var (
	re *regexp.Regexp // = regexp.MustCompile("[a-z]{2}")
)

func init() {
	re = regexp.MustCompile("[a-z]{2}")
}

func findAllString() []string {
	// re := regexp.MustCompile("[a-z]{2}")
	re2 := re.Copy()
	return re2.FindAllString("hoge", -1)
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			fmt.Println(findAllString())
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Hello, playground")
}
