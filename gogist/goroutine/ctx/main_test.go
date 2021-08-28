package ctx

import (
	"context"
	"testing"
	"time"
)

func TestTimeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	go monitor(ctx)

	select {
	case <-ctx.Done():
		t.Log(ctx.Err())
	case <-time.After(10 * time.Second):
		// 等待协程执行
		t.Log("timeout")
	}
}
