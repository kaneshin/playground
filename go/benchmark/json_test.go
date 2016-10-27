package benchmark

import (
	"bytes"
	"encoding/json"
	"testing"
)

func BenchmarkJSONDecode(b *testing.B) {
	const jsonString = `{"foo":true,"bar":1,"hoge":"hogehoge"}`

	b.Run("Unmarshal", func(b *testing.B) {
		var d map[string]interface{}
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			json.Unmarshal([]byte(jsonString), &d)
		}
	})

	b.Run("Decoder", func(b *testing.B) {
		var d map[string]interface{}
		buf := bytes.NewBufferString(jsonString)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			json.NewDecoder(buf).Decode(&d)
		}
	})
}
