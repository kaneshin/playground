package main

import (
	"fmt"
	"io/ioutil"
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
		b, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(b))
	} else {
		fmt.Println("no named pipe:", m.String())
	}
}
