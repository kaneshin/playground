package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

//+TP
type Tag struct {
	Name string
}

//+VR
var fooTag = Tag{
	Name: "foo",
}

//+VR
var barTag = Tag{
	Name: "bar",
}

//+VR
var (
	fooTag1 = Tag{
		Name: "foo",
	}

	barTag1 = Tag{
		Name: "bar",
	}
)

func main() {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "./main.go", nil, parser.ParseComments)
	if err != nil {
		return
	}

	for _, decl := range f.Decls {
		d, ok := decl.(*ast.GenDecl)
		if !ok {
			continue
		}

		doc := d.Doc
		if doc == nil {
			continue
		}

		for _, comment := range doc.List {
			switch comment.Text {
			case "//+VR":
				fmt.Printf("-- //+VR\n")
				for _, spec := range d.Specs {
					fmt.Printf("%#v\n", spec)
				}

			case "//+TP":
				fmt.Printf("-- //+TP\n")
				for _, spec := range d.Specs {
					fmt.Printf("%#v\n", spec)
				}

			}
		}

	}
}
