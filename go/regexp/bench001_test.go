package regexp

import (
	"testing"
)

const (
	sample = `こんにちは！こんにちは！
あ　し　た　ー
あ し た -
おお

oo

`
	expected = `こんにちは！こんにちは！あしたーあした-おおoo`
)

func TestWhitespaceTrim(t *testing.T) {
	var (
		result string
	)

	result = ""
	result = WhitespaceTrimByStringsFields(sample)
	if result != expected {
		t.Fatalf("Expected %v, but %v:", expected, result)
	}

	result = ""
	result = WhitespaceTrimByRegexp(sample)
	if result != expected {
		t.Fatalf("Expected %v, but %v:", expected, result)
	}
}

func BenchmarkStringsFields(b *testing.B) {
	for n := 0; n < b.N; n++ {
		WhitespaceTrimByStringsFields(sample)
	}
}

func BenchmarkRegexp(b *testing.B) {
	for n := 0; n < b.N; n++ {
		WhitespaceTrimByRegexp(sample)
	}
}
