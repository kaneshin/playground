package regexp

import (
	"regexp"
	"strings"
)

var (
	_ = strings.Join
	_ = regexp.Compile

	trim = regexp.MustCompile("[\r|\n| |　]")
)

func WhitespaceTrimByStringsFields(msg string) string {
	return strings.Join(strings.Fields(msg), "")
}

func WhitespaceTrimByRegexp(msg string) string {
	return trim.ReplaceAllString(msg, "")
}
