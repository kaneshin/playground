# PROTO_PATH=/path/to/github.com/google/protobuf/src

protoc:
	@go get -u github.com/golang/protobuf/protoc-gen-go

build: protoc
	@protoc --proto_path=$$PROTO_PATH --proto_path=. --go_out=. *.proto
