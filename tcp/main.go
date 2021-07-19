package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "192.168.115.48:8000")
	if nil != err {
		log.Println(err)
	}

	fmt.Println(conn.LocalAddr().String())
	fmt.Println(conn.RemoteAddr().String())

	go func() {
		data := make([]byte, 4096)

		for {
			n, err := conn.Read(data)
			if err != nil {
				log.Println("hihihi")
				log.Println(err)
				return
			}

			log.Println("Server send : " + string(data[:n]))
			time.Sleep(time.Duration(3) * time.Second)
		}
	}()

	for {
		var s string
		fmt.Scanln(&s)
		conn.Write([]byte(s))
		time.Sleep(time.Duration(3) * time.Second)
	}
}
