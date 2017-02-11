package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func printCsv(r io.Reader) error {
	records, err := csv.NewReader(r).ReadAll()
	if err != nil {
		return err
	}

	for _, line := range records {
		fmt.Println(line)
	}

	return nil

}

func main() {
	f, err := os.Open("test.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	printCsv(f)
}
