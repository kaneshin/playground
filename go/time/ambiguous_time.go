package main

import (
	"encoding/json"
	"fmt"
	"time"
)

const value = `{
  "fixedAt": "2016-03-23T11:28:41+09:00",
  "registeredAt": false
}`

type (
	AmbiguousTime struct {
		Raw interface{}
	}
)

func (t AmbiguousTime) Time() time.Time {
	if v, ok := t.Raw.(string); ok {
		u, _ := time.Parse(time.RFC3339, v)
		return u
	}
	return time.Time{}
}

func (t AmbiguousTime) After(u time.Time) bool {
	v := t.Time()
	if v.IsZero() {
		return false
	}
	return v.After(u)
}

func (t AmbiguousTime) Before(u time.Time) bool {
	v := t.Time()
	if v.IsZero() {
		return false
	}
	return v.Before(u)
}

var v = struct {
	FixedAt      interface{} `json:"fixedAt"`
	RegisteredAt interface{} `json:"registeredAt"`
}{}

func main() {
	fmt.Println(json.Unmarshal([]byte(value), &v))
	fmt.Println(AmbiguousTime{v.FixedAt}.Time())
	fmt.Println(AmbiguousTime{v.RegisteredAt}.Time())
	fmt.Println(AmbiguousTime{v.FixedAt}.Before(time.Now()))
	fmt.Println(AmbiguousTime{v.RegisteredAt}.Before(time.Now()))
	fmt.Println(AmbiguousTime{v.FixedAt}.After(time.Now()))
	fmt.Println(AmbiguousTime{v.RegisteredAt}.After(time.Now()))
}

