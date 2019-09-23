package timeout

import (
	"errors"
	"time"
)

func timeoutWrap(f func()) error {
	ch := make(chan int)
	go func() {
		f()
		ch <- 1
	}()

	select {
	case <-ch:
		return nil
	case <-time.After(3 * time.Second):
		return errors.New("timeout")
	}
}

func add(a, b int) int {
	time.Sleep(5 * time.Second)
	return a + b
}
