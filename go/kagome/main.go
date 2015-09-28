package main

import (
	"fmt"
	"strings"

	"github.com/ikawaha/kagome"
)

// main ...
func main() {
	t := kagome.NewTokenizer()
	tokens := t.Tokenize("寿司が食べたい。")
	for _, token := range tokens {
		if token.Class == kagome.DUMMY {
			// BOS: Begin Of Sentence, EOS: End Of Sentence.
			// fmt.Printf("%s\n", token.Surface)
			continue
		}
		features := strings.Join(token.Features(), ",")
		fmt.Printf("%s\t%v\n", token.Surface, features)
	}
}
