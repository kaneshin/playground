package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func init() {
	flag.Parse()
}

func main() {
	var name string
	if args := flag.Args(); len(args) > 0 {
		name = args[0]
	}

	var r io.Reader
	switch name {
	case "", "-":
		r = os.Stdin
	default:
		f, err := os.Open(name)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		r = f
	}

	b, err := ioutil.ReadAll(r)
	if err != nil {
		panic(err)
	}
	fmt.Print(string(b))
}
