package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	buf := bytes.NewBufferString(`{"id":1}`)
	u := new(User)
	if err := json.NewDecoder(buf).Decode(u); err != nil {
		panic(err)
	}

	fmt.Printf("%#v", *u)
}
