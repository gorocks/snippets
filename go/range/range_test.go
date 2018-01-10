package range_test

import (
	"reflect"
	"testing"
)

func TestRangeLoop(t *testing.T) {
	v0 := []int{1, 2, 3}
	for i := range v0 {
		v0 = append(v0, i)
	}
	v1 := []int{1, 2, 3, 0, 1, 2}
	if !reflect.DeepEqual(v0, v1) {
		t.Errorf("v want: %v, but is: %v", v0, v1)
	}

	m := map[string]int{
		"a": 1,
		"b": 2,
	}

	j := 0
	for _, v := range m {
		m["c"] = v + 1
		j++
	}

	t.Log(j)
	// here j maybe 2 or 3
}
