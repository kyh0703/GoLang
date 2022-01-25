package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/opentracing/opentracing-go/log"
	"golang.org/x/sync/errgroup"
)

func test1(ctx context.Context) error {
	timer := time.NewTimer(time.Second * 5)
	select {
	case <-timer.C:
		fmt.Println("test1 timer")
		return nil
	case <-ctx.Done():
		fmt.Println("test1 cancel")
		return ctx.Err()
	}
}

func test2(ctx context.Context) error {
	timer := time.NewTimer(time.Second * 2)
	select {
	case <-timer.C:
		// TEST2: go routine 내부에서 에러 발생
		// TEST2: context=> test1이 종료되지않는다.
		fmt.Println("test2 timer")
		fmt.Println("Occur error test2")
		return errors.New("error test2")
	case <-ctx.Done():
		fmt.Println("test2 cancel")
		return ctx.Err()
	}
}

func main() {
	// ctx, _ := context.WithCancel(context.Background())
	// errGrp, _ := errgroup.WithContext(ctx)

	ctx, _ := context.WithCancel(context.Background())
	errGrp, errCtx := errgroup.WithContext(ctx)

	errGrp.Go(func() error { return test1(errCtx) })
	errGrp.Go(func() error { return test2(errCtx) })

	// TEST1: 넘겨준 context cancel 발생시 처리
	// TEST1 RESULT: go routine 2개 종료

	// timer := time.NewTimer(time.Second * 2)
	// <-timer.C
	// cancel()

	if err := errGrp.Wait(); err != nil {
		log.Error(err)
	}
}
