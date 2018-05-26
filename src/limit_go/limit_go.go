package limit_go

import (
	"sync"
)

type GoLimit struct {
	ch    chan int
	wg    sync.WaitGroup
	count int
}

func NewGoLimit(x int) *GoLimit {
	return &GoLimit{
		make(chan int, x),
		sync.WaitGroup{},
		0,
	}
}

func (gl *GoLimit) LimitCall(cb func()) {
	gl.ch <- 1
	gl.wg.Add(1)
	gl.count++
	go gl.LimitWork(cb)
}

func (gl *GoLimit) LimitWork(cb func()) {
	defer func() {
		<-gl.ch
		gl.wg.Done()
		gl.count--
	}()

	cb()
}
