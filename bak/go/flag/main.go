package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	fmt.Printf("%v\n", os.Args[1:])
	defFoo := flag.String("foo", "", "")
	flag.Parse()
	fmt.Printf("%v\n", flag.Args())
	fmt.Printf("%s\n", *defFoo)

	flg := flag.NewFlagSet("flag", flag.PanicOnError)
	mkFoo := flg.String("foo", "", "")
	flg.Parse(os.Args[1:])
	fmt.Printf("%v\n", flg.Args())
	fmt.Printf("%s\n", *mkFoo)
}
