package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
)

const (
	// SOI is a marker of the start of image.
	SOI = 0xd8

	// EOI is a marker of the end of image.
	EOI = 0xd9

	// SOS is a marker of the start of scan.
	SOS = 0xda

	// SOF0 is a marker of the start of frame.
	// Huffman coding, Baseline JPEG
	SOF0 = 0xc0

	// APP0 is a marker of the application.
	APP0 = 0xe0

	// APP1 is a marker of the application.
	APP1 = 0xe1
)

func validate(b []byte) bool {
	if len(b) < 2 {
		return false
	}
	return b[0] == 0xff && b[1] == SOI
}

func has(n byte, b []byte) bool {
	for i, c := range b {
		if c == 0xff {
			switch b[i+1] {
			case n:
				return true

			case SOS:
				return false

			}
		}
	}
	return false
}

func hasJFIF(b []byte) bool {
	return has(APP0, b)
}

func hasExif(b []byte) bool {
	return has(APP1, b)
}

func hasSOF0(b []byte) bool {
	return has(SOF0, b)
}

func main() {
	flag.Parse()

	if len(flag.Args()) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	name := flag.Args()[0]
	data, err := ioutil.ReadFile(name)
	if err != nil {
		log.Fatalf("error %v", err)
	}

	if len(data) == 0 || !validate(data) {
		log.Fatalf("invalid jpeg file  %v", name)
	}

	if hasJFIF(data) {
		log.Printf("data has JFIF Field")
	}

	if hasExif(data) {
		log.Printf("data has Exif Field")
	}

	if hasSOF0(data) {
		log.Printf("data has SOF0 Field")
	}
}
