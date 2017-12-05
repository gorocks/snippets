package reflect_test

import (
	"reflect"
	"testing"
)

func TestValueString(t *testing.T) {
	tests := []struct {
		name string
		arg  interface{}
		want string
	}{
		{"string value", "abcd", "abcd"},
		{"int value", 5, "<int Value>"},
		{"int value", 0, "<int Value>"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reflect.ValueOf(tt.arg).String(); got != tt.want {
				t.Errorf("reflect.ValueOf().String() = %v, want %v", got, tt.want)
			}
		})
	}
}
