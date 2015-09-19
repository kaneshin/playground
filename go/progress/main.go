package main

import (
	"fmt"
	"time"
)

func progress() {
	var i int8 = 0
	for {
		switch i % 4 {
		case 0:
			fmt.Print("|")
		case 1:
			fmt.Print("/")
		case 2:
			fmt.Print("-")
		case 3:
			fmt.Print("\\")
			i = -1
		}
		i++
		time.Sleep(100 * time.Millisecond)
		fmt.Print("\r")
	}
}

func main() {
	go progress()
	time.Sleep(10 * time.Second)
}
