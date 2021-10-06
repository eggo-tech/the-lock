package spin

// Spin lock
type Spin int32

// Lock ...
func (l *Spin) Lock() {
	lock((*int32)(l), 0, 1)
}

// Unlock ...
func (l *Spin) Unlock() {
	unlock((*int32)(l), 0)
}

func lock(ptr *int32, o, n int32)
func unlock(ptr *int32, n int32)
