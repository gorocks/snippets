package syncs_test

import (
	"sync"
	"testing"

	"github.com/gorocks/snippets/go/syncs"
)

func BenchmarkSyncMapStore(b *testing.B) {
	var bm sync.Map
	for n := 0; n < b.N; n++ {
		bm.Store("foo", "bar")
	}
}
func BenchmarkMapSet(b *testing.B) {
	var m syncs.Map
	for n := 0; n < b.N; n++ {
		m.Set("foo", "bar")
	}
}

// goos: darwin
// goarch: amd64
// pkg: github.com/gorocks/snippets/go/syncs
// BenchmarkSyncMapStore-8   	20000000	        95.9 ns/op	      16 B/op	       1 allocs/op
// BenchmarkMapSet-8         	20000000	        64.9 ns/op	       0 B/op	       0 allocs/op
