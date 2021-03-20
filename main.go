package main

import (
	"flag"
	"github.com/eggo-tech/the-lock/atom"
	"github.com/eggo-tech/the-lock/spin"
	"time"
)

type Locker interface {
	Lock()
	Unlock()
}

func main() {
	var p bool
	flag.BoolVar(&p, "pause", false, "with pause instruction")
	flag.Parse()
	var l Locker
	if p {
		l = new(spin.Spin)
	} else {
		l = new(atom.Spin)
	}
	var n int
	for i := 0; i < 2; i++ {
		go routine(i, &n, l, 500*time.Millisecond)
	}
	select {}
}

func routine(i int, v *int, l Locker, d time.Duration) {
	for {
		func() {
			l.Lock()
			defer l.Unlock()
			*v++
			println(*v, i)
			time.Sleep(d)
		}()
	}
}
