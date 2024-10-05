package closing

import (
	"context"
	"fmt"
	"time"
)

func RunGoroutineClosingExample() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := Execute(ctx); err != nil {
		fmt.Printf("Err: %s\n", err.Error())
	}
}

func Execute(ctx context.Context) error {
	firstResult := make(chan struct{})
	secondResult := make(chan struct{})

	go func() {
		time.Sleep(10000 * time.Millisecond)
		firstResult <- struct{}{}
	}()

	go func() {
		time.Sleep(5000 * time.Millisecond)
		secondResult <- struct{}{}
	}()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Timeout executed")
			return ctx.Err()
		case <-firstResult:
			fmt.Println("Got first goroutine result")
			return nil
		case <-secondResult:
			fmt.Println("Got second goroutine result")
			return nil
		}
	}
}
