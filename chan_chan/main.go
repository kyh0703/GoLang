package main

import (
	"context"
	"log"
	"sync"
	"time"
)

var wg sync.WaitGroup

func consumer(ctx context.Context, reqCh chan chan string) {
	defer wg.Done()

	log.Print("consumer")
	repCh := make(chan string)
	defer close(repCh)

	// Init Response wait
	time.Sleep(time.Second)

	log.Print("consumer make chan")
	reqCh <- repCh

	for {
		select {
		case msg := <-repCh:
			log.Print(msg)
		case <-ctx.Done():
			log.Print("close consumer")
			return
		}
	}
}

func producer(ctx context.Context, reqCh chan chan string) {
	defer wg.Done()
	log.Print("producer wait chan")
	repCh := <-reqCh
	repCh <- "hihihihi"
}

func main() {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*3)
	reqCh := make(chan chan string)

	wg.Add(1)
	go consumer(ctx, reqCh)
	wg.Add(1)
	go producer(ctx, reqCh)

	wg.Wait()
	close(reqCh)
}
