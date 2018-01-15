package map_test

import (
	"testing"
)

func TestMap(t *testing.T) {
	t.Log((len(map[interface{}]int{
		new(int):      1,
		new(int):      2,
		new(struct{}): 3,
		new(struct{}): 4,
	}))) // you may want 4 but it's 3 in go 1.9.2.
}
