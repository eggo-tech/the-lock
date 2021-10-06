package atom

import (
	"sync/atomic"
)

// Spin lock
type Spin [2]int32

// Lock ...
func (l *Spin) Lock() {
	for !atomic.CompareAndSwapInt32(&l[0], 0, 1) {
		if c := l[1]; c > 0 {
			pause(int(c))
		}
	}
}

// Unlock ...
func (l *Spin) Unlock() {
	atomic.StoreInt32(&l[0], 0)
}

func pause(cnt int)
