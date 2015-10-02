package bench

import (
	"bytes"
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
