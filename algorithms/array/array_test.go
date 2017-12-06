package array_test

import (
	"reflect"
	"testing"

	"github.com/gorocks/snippets/algorithms/array"
)

func TestMergeTwoArray(t *testing.T) {
	type args struct {
		a []int
		b []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Merge of '[1,2,3,4,5]' and '[0,6,7,8] is '[0,1,2,3,4,5,6,7,8]'", args{[]int{1, 2, 3, 4, 5}, []int{0, 6, 7, 8}}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8}},
		{"Merge of '[1,2]' and '[0,1,2,3] is '[0,1,1,2,2,3]'", args{[]int{1, 2}, []int{0, 1, 2, 3}}, []int{0, 1, 1, 2, 2, 3}},
		{"Merge of '[4,5]' and '[0,1,2,3] is '[0,1,2,3,4,5]'", args{[]int{4, 5}, []int{0, 1, 2, 3}}, []int{0, 1, 2, 3, 4, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := array.MergeTwoArray(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeTwoArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
