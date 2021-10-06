package atom

import (
	"sync/atomic"
)

type Spin int32

var Loop int

func (l *Spin) Lock() {
	for !atomic.CompareAndSwapInt32((*int32)(l), 0, 1) {
		if Loop > 0 {
			pause(Loop)
		}
	}
}

func (l *Spin) Unlock() {
	atomic.StoreInt32((*int32)(l), 0)
}

func pause(cnt int)
