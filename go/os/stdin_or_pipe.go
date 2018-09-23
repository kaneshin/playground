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
	fi, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	var r io.Reader
	if fi.Mode()&os.ModeNamedPipe == os.ModeNamedPipe {
		// mode named pipe
		r = os.Stdin
	} else {
		args := flag.Args()
		switch len(args) {
		case 0:
		case 1:

		}
		if len(args) == 0 {
			r = os.Stdin
		} else {
			name := args[0]
			if name == "-" {
				r = os.Stdin
			} else {
				f, err := os.Open(name)
				if err != nil {
					panic(err)
				}
				defer f.Close()
				r = f
			}
		}
	}

	b, err := ioutil.ReadAll(r)
	if err != nil {
		panic(err)
	}
	fmt.Print(string(b))
}
