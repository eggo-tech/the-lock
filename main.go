package main

import (
	"flag"
	"github.com/eggo-tech/the-lock/atom"
	"github.com/eggo-tech/the-lock/spin"
	"sync"
)

type Locker interface {
	Lock()
	Unlock()
}

type Waiter interface {
	Add(n int)
	Wait()
}

func main() {
	var p bool
	flag.BoolVar(&p, "pause", false, "with pause instruction")
	flag.IntVar(&atom.Loop, "loop", 0, "pause loop times")
	flag.Parse()
	var l Locker
	if p {
		l = new(spin.Spin)
	} else {
		l = new(atom.Spin)
	}
	var w Waiter
	w = new(sync.WaitGroup)
	var n int
	for i := 0; i < 2; i++ {
		w.Add(1)
		d := 1
		if i%2 != 0 {
			d = -1
		}
		go routine(i, &n, l, w, 100000000, d)
	}
	w.Wait()
	println(n)
}

func routine(i int, v *int, l Locker, w Waiter, c, d int) {
	defer w.Add(-1)
	for t := 0; t < c; t++ {
		func() {
			l.Lock()
			defer l.Unlock()
			*v += d
		}()
	}
}
