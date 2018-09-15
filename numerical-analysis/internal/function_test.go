package internal

import (
	"reflect"
	"testing"
)

func TestFunction(t *testing.T) {
	var f Function

	tests := map[string]struct {
		f          interface{}
		assignable bool
	}{
		"F1": {F1, true},
		"F2": {F2, true},
		"NG": {100, false},
	}

	for k, tt := range tests {
		tt := tt
		t.Run(k, func(t *testing.T) {
			ft := reflect.TypeOf(f)
			tft := reflect.TypeOf(tt.f)
			ok := tft.AssignableTo(ft)
			switch {
			case !tt.assignable && ok:
				t.Fatalf("This is assignable `Function' type but not expected")
			case tt.assignable && !ok:
				t.Fatalf("This is not able to be `Function' type assertion")
			}
		})
	}
}
