package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("test.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	records, err := csv.NewReader(f).ReadAll()
	if err != nil {
		panic(err)
	}

	for _, line := range records {
		fmt.Println(line)
	}
}
