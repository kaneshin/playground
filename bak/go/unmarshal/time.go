package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type ExTime struct {
	time.Time
}

func now() ExTime {
	return ExTime{time.Now()}
}

type Transaction struct {
	Date1 ExTime `json:"date1"`
	Date2 ExTime `json:"date2,integer"`
	Date3 ExTime `json:"date3,string"`
}

func main() {
	txn := Transaction{
		Date1: now(),
		Date2: now(),
		Date3: now(),
	}

	out, _ := json.Marshal(txn)
	fmt.Println(string(out))
}
