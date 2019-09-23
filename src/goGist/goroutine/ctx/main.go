package ctx

import (
	"context"
	"fmt"
	"time"
)

func monitor(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("done")
			return
		default:
			fmt.Println("running")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
