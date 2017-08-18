package json

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func Test_merge(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected map[string]interface{}
	}{
		{
			name:  "test01",
			input: User{1, "test01", "test01@example.com"},
			expected: map[string]interface{}{
				"meta":  meta,
				"id":    1.0,
				"name":  "test01",
				"email": "test01@example.com",
			},
		},
		{
			name: "test02",
			input: map[string]interface{}{
				"instances": []User{
					User{1, "test01", "test01@example.com"},
					User{2, "test02", "test02@example.com"},
				},
			},
			expected: map[string]interface{}{
				"meta": meta,
				"instances": []interface{}{
					map[string]interface{}{
						"id":    1.0,
						"name":  "test01",
						"email": "test01@example.com",
					},
					map[string]interface{}{
						"id":    2.0,
						"name":  "test02",
						"email": "test02@example.com",
					},
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert := assert.New(t)
			assert.EqualValues(tt.expected, merge_v1(tt.input))
			assert.EqualValues(tt.expected, merge_v2(tt.input))
			assert.EqualValues(tt.expected, merge_v3(tt.input))
		})
	}
}

var data = map[string]interface{}{
	"instances": []User{
		User{1, "test01", "test01@example.com"},
		User{2, "test02", "test02@example.com"},
		User{3, "test03", "test03@example.com"},
		User{4, "test04", "test04@example.com"},
		User{5, "test05", "test05@example.com"},
		User{6, "test06", "test06@example.com"},
		User{7, "test07", "test07@example.com"},
		User{8, "test08", "test08@example.com"},
		User{9, "test09", "test09@example.com"},
		User{10, "test10", "test10@example.com"},
		User{11, "test11", "test11@example.com"},
		User{12, "test12", "test12@example.com"},
		User{13, "test13", "test13@example.com"},
		User{14, "test14", "test14@example.com"},
		User{15, "test15", "test15@example.com"},
		User{16, "test16", "test16@example.com"},
		User{17, "test17", "test17@example.com"},
		User{18, "test18", "test18@example.com"},
		User{19, "test19", "test19@example.com"},
		User{20, "test20", "test20@example.com"},
		User{21, "test21", "test21@example.com"},
		User{22, "test22", "test22@example.com"},
		User{23, "test23", "test23@example.com"},
		User{24, "test24", "test24@example.com"},
		User{25, "test25", "test25@example.com"},
		User{26, "test26", "test26@example.com"},
		User{27, "test27", "test27@example.com"},
		User{28, "test28", "test28@example.com"},
		User{29, "test29", "test29@example.com"},
		User{30, "test30", "test30@example.com"},
		User{31, "test31", "test31@example.com"},
		User{32, "test32", "test32@example.com"},
		User{33, "test33", "test33@example.com"},
		User{34, "test34", "test34@example.com"},
		User{35, "test35", "test35@example.com"},
		User{36, "test36", "test36@example.com"},
		User{37, "test37", "test37@example.com"},
		User{38, "test38", "test38@example.com"},
		User{39, "test39", "test39@example.com"},
		User{40, "test40", "test40@example.com"},
		User{41, "test41", "test41@example.com"},
		User{42, "test42", "test42@example.com"},
		User{43, "test43", "test43@example.com"},
		User{44, "test44", "test44@example.com"},
		User{45, "test45", "test45@example.com"},
		User{46, "test46", "test46@example.com"},
		User{47, "test47", "test47@example.com"},
		User{48, "test48", "test48@example.com"},
		User{49, "test49", "test49@example.com"},
		User{50, "test50", "test50@example.com"},
	},
}

var result interface{}

func Benchmark_merge(b *testing.B) {
	b.Run("merge_v1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			result = merge_v1(data)
		}
	})

	b.Run("merge_v2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			result = merge_v2(data)
		}
	})

	b.Run("merge_v3", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			result = merge_v3(data)
		}
	})
}
