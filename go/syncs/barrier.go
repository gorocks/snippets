package syncs

import (
	"sync"
)

var flag = uint(1) << 31

// Barrier implements POSIX pthread barrier.
type Barrier struct {
	count uint
	total uint
	cond  *sync.Cond
}

// NewBarrier returns a new Barrier with uint count.
func NewBarrier(count uint) *Barrier {
	return &Barrier{
		count: count,
		total: flag,
		cond:  sync.NewCond(new(sync.Mutex)),
	}
}

// Wait synchronize at a Barrier.
func (b *Barrier) Wait() {
	b.cond.L.Lock()
	for b.total > flag {
		b.cond.Wait()
	}
	if b.total == flag {
		b.total = 0
	}
	b.total++
	if b.total == b.count {
		b.total += flag - 1
		b.cond.Broadcast()
		b.cond.L.Unlock()
	} else {
		for b.total < flag {
			b.cond.Wait()
		}
		b.total--
		if b.total == flag {
			b.cond.Broadcast()
		}
		b.cond.L.Unlock()
	}
}
