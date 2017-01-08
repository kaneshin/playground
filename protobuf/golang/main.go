package main

import (
	"log"
	"net/http"

	proto1 "github.com/golang/protobuf/proto"
	proto "github.com/kaneshin/playground/protobuf"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		param := proto.Parameter{
			Key:   "name",
			Value: "#ff0000",
		}

		data, err := proto1.Marshal(&param)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		w.Write(data)

		param1 := &proto.Parameter{}
		if err := proto1.Unmarshal(data, param1); err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		log.Print(param1.String())
	})
	http.ListenAndServe(":8080", nil)
}
