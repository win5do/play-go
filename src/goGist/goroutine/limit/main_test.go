package limit

import (
	"testing"
	"time"
)

func work() {
	time.Sleep(1 * time.Second)
}

func TestWg(t *testing.T) {
	wl := NewGoLimit(10)
	for i := 0; i <= 30; i++ {
		t.Log(wl.count)
		wl.LimitCall(work)
	}
	wl.wg.Wait()
}
