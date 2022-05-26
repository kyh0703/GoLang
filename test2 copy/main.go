package main

import (
	"context"
	"fmt"
	"time"
)

const MaxEvent = 30

type Result struct {
	eventType uint16
	err       error
	time      time.Time
}

type CompatibleFlow struct {
	errCh  chan error
	procCh chan int
}

func NewCompatibleFlow() *CompatibleFlow {
	return &CompatibleFlow{
		errCh:  make(chan error),
		procCh: make(chan int),
	}
}

func (f *CompatibleFlow) Err() <-chan error {
	return f.errCh
}

func (f *CompatibleFlow) Process() <-chan int {
	return f.procCh
}

func (f *CompatibleFlow) Validate() error {
	return nil
}

func (f *CompatibleFlow) DoRoute(ctx context.Context) {
	go func() {
		defer close(f.errCh)
		defer close(f.procCh)
		if err := f.Validate(); err != nil {
			f.errCh <- err
			return
		}
		fmt.Println("process input")
		// f.errCh <- errors.New("erer")
		f.procCh <- 3
	}()
}

func main() {
	route := NewCompatibleFlow()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	route.DoRoute(ctx)
	for {
		select {
		case <-ctx.Done():
			return
		case err := <-route.Err():
			fmt.Println("Error", err)
			return
		case proc := <-route.Process():
			fmt.Println("process done")
			fmt.Println(proc)
		}
	}
}
