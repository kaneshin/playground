package main

import (
	"fmt"
	"go/build"
	"path/filepath"
	"reflect"
)

func main() {
	gopaths := filepath.SplitList(build.Default.GOPATH)
	fmt.Println(gopaths[0])
	fmt.Println(reflect.TypeOf(gopaths[0]))
}
