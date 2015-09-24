package main

import (
	"fmt"

	"github.com/microcosm-cc/bluemonday"
)

// main ...
func main() {
	p := bluemonday.UGCPolicy()
	html := p.Sanitize(
		`<a onblur="alert(secret)" href="http://www.google.com">Google</a>`,
	)

	// Output:
	// <a href="http://www.google.com" rel="nofollow">Google</a>
	fmt.Println(html)
}
