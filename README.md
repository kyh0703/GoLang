Go chan chan이란?

소비자가 컨슈머에게 받을 준비가 된 후에 채널을 전달히면 공급자는 소비지가 받을 준비가 된 것을 확인하고 메시지를 보낸다. 

package main

import "fmt"
import "time"

func main() {
     // 1. chan chan을 생성 후 고루틴에 전달한다
     requestChan := make(chan chan string)

     go goroutineC(requestChan)
     go goroutineD(requestChan)

     time.Sleep(time.Second)

}

func goroutineC(requestChan chan chan string) {
     // 소비자는 받을 채널을 생성후에 생성자에게 채널을 전달한다
     responseChan := make(chan string)

     requestChan <- responseChan

     response := <-responseChan

     fmt.Printf("Response: %v\n", response)

}

func goroutineD(requestChan chan chan string) {
     // 생산자는 소비자에게서 전송 받을 채널이 들어오면 메시지를 전달한다.
     responseChan := <-requestChan

     responseChan <- "wassup!"
}
