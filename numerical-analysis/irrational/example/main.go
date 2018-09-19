package main

import (
	"fmt"

	"github.com/kaneshin/playground/numerical-analysis/irrational"
	_ "github.com/kaneshin/playground/numerical-analysis/irrational/leibniz"
)

func main() {
	fmt.Println(irrational.Pi)
}
