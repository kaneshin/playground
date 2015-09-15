package main

import (
	"fmt"
	"github.com/kaneshin/gopack/exception"
)

// main ...
func main() {
	fmt.Println("Hello world")

	exception.Try(func() {
		fmt.Println("There is no problem.")
	}).Catch(func(e exception.Exception) {
		fmt.Println("Catch exception.")
	}).Finally(func() {
		fmt.Println("Do finally.")
	})

	exception.Try(func() {
		fmt.Println("There is a problem.")
		panic("Panic happens.")
	}).Catch(func(e exception.Exception) {
		fmt.Println("Catch exception.")
	}).Finally(func() {
		fmt.Println("Do finally.")
	})

}
