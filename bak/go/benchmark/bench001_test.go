package bench

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

func BenchmarkConcatJoin(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ids := []string{"foo", "bar", "baz"}
		func(ids []string) {
			for i, id := range ids {
				ids[i] = fmt.Sprintf(`'%s'`, id)
			}
		}(ids)
		_ = strings.Join(ids, ",")
	}
}

func BenchmarkConcatBuffer(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ids := []string{"foo", "bar", "baz"}
		var byteWhereInIDs bytes.Buffer
		for _, id := range ids {
			byteWhereInIDs.WriteString("'")
			byteWhereInIDs.WriteString(id)
			byteWhereInIDs.WriteString("'")
			byteWhereInIDs.WriteString(",")
		}
		whereInIDs := byteWhereInIDs.String()
		_ = whereInIDs[0:(len(whereInIDs) - 1)]
	}
}

var jsonData = []byte(`{"foo":true,"bar":1,"hoge":"hogehoge"}`)
var data map[string]interface{}

func BenchmarkUnmarshal(b *testing.B) {
	for n := 0; n < b.N; n++ {
		json.Unmarshal(jsonData, &data)
	}
}

var buf bytes.Buffer

func init() {
	buf.Write([]byte(`{"foo":true,"bar":1,"hoge":"hogehoge"}`))
}

func BenchmarkDecoder(b *testing.B) {
	for n := 0; n < b.N; n++ {
		json.NewDecoder(&buf).Decode(&data)
	}
}
