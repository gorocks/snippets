package syncs_test

import (
	"fmt"
	"testing"

	"github.com/gorocks/snippets/go/syncs"
)

func TestBarrierWait(t *testing.T) {
	b := syncs.NewBarrier(5)
	for i := 0; i < 4; i++ {
		go func(i int) {
			fmt.Println("I am ", i)
			b.Wait()
			fmt.Println("I am done ...")
		}(i)
	}
	b.Wait()
	fmt.Println("all are done ...")
	// I am  3
	// I am  2
	// I am  1
	// I am  0
	// I am done ...
	// I am done ...
	// I am done ...
	// all are done ...
	// I am done ...
}
