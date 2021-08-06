package main

import (
	"fmt"
	"net/http"
	"time"
)

func ctxTest(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	fmt.Println("processing request")

	select {
	case <-time.After(10 * time.Second):
		w.Write([]byte("request processed"))
		fmt.Println("hihihihihi")
	case <-ctx.Done():
		fmt.Println("request cancelled")
	}
}

func main() {
	http.ListenAndServe("localhost:8000", http.HandlerFunc(ctxTest))
}
