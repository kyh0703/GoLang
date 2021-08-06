package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func operation1(ctx context.Context) error {
	time.Sleep(5 * time.Second)
	return errors.New("failed")
}

func operation2(ctx context.Context) {
	select {
	case <-time.After(10 * time.Second):
		fmt.Println("done")
	case <-ctx.Done():
		fmt.Println("operation2 cancelled")
	}
}

func main() {
	// Create a new context
	ctx := context.TODO()
	ctx2, cancel := context.WithCancel(ctx)
	go func() {
		err := operation1(ctx2)
		if err != nil {
			cancel()
		}
	}()
	operation2(ctx2)
}
