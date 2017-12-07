package tree

import (
	"reflect"
	"testing"
)

func TestTreeInOrder(t *testing.T) {
	tree := &Tree{
		Value: 4,
		Left: &Tree{
			Value: 2,
			Left:  &Tree{Value: 1},
			Right: &Tree{Value: 3},
		},
		Right: &Tree{Value: 5}}
	if want, got := []int{1, 2, 3, 4, 5}, tree.InOrder(); !reflect.DeepEqual(want, got) {
		t.Errorf("Tree.InOrder() = %v, want %v", got, want)
	}
}

func TestNewInOrderTree(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	if want, got := a, NewInOrderTree(a).InOrder(); !reflect.DeepEqual(want, got) {
		t.Errorf("NewInOrderTree() = %v, want %v", got, want)
	}
}
