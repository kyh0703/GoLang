// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )

// func factoryRand(count int) (out chan int) {
// 	out = make(chan int)
// 	go func() {
// 		for i := 0; i < count; i++ {
// 			out <- rand.Intn(100)
// 		}
// 		close(out)
// 	}()
// 	return out
// }

// func printNumber(in <-chan int) <-chan int {
// 	out := make(chan int)
// 	go func() {
// 		for n := range in {
// 			fmt.Println("print: ", n)
// 			out <- n
// 		}
// 		close(out)
// 	}()
// 	return out
// }

// func squareNumber(in <-chan int) <-chan int {
// 	out := make(chan int)
// 	go func() {
// 		for n := range in {
// 			out <- n * n
// 		}
// 		close(out)
// 	}()
// 	return out
// }

// // func main() {
// // 	// 작업시작시간 기록
// // 	start := time.Now()

// // 	c0 := factoryRand(100000)
// // 	c1 := printNumber(c0)
// // 	c2 := squareNumber(c1)

// // 	// 최종적으로 더하는 작업을 메인 프로세스에서 진행합니다
// // 	sum := 0
// // 	for n := range c2 {
// // 		sum += n
// // 	}
// // 	fmt.Printf("Total Sum of Squares: %d\n", sum)

// // 	// 작업 종료 후 시간기록
// // 	elapsed := time.Since(start)
// // 	fmt.Println("작업소요시간: ", elapsed)
// // }

// func main() {
// 	// 작업시작시간 기록
// 	start := time.Now()

// 	sum := 0
// 	for n := range squareNumber(printNumber(factoryRand(100000))) {
// 		sum += n
// 	}
// 	fmt.Printf("Total Sum of Squares: %d\n", sum)

// 	// 작업 종료 후 시간기록
// 	elapsed := time.Since(start)
// 	fmt.Println("작업소요시간: ", elapsed)
// }
