package main

import (
	"log"
	"net/http"

	"github.com/golang/protobuf/proto"
	"github.com/kaneshin/playground/protobuf"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		user := &protobuf.User{
			Id:        1,
			FirstName: "first name",
			LastName:  "last name",
		}

		data, err := proto.Marshal(user)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		w.Write(data)

		other := &protobuf.User{}
		if err := proto.Unmarshal(data, other); err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		log.Print(other.String())
	})
	http.ListenAndServe(":8080", nil)
}
