package main

import (
	"fmt"
	"os"
)

func main() {
	fi, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	m := fi.Mode()
	if m&os.ModeNamedPipe == os.ModeNamedPipe {
		fmt.Println("named pipe:", m.String())
	} else {
		fmt.Println("no named pipe:", m.String())
	}
}
