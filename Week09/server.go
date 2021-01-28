package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	fmt.Println("Starting the server ...")
	// 创建 listener
	listener, err := net.Listen("tcp", "localhost:50000")
	if err != nil {
		fmt.Println("Error listening", err.Error())
		return //终止程序
	}
	// 监听并接受来自客户端的连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting", err.Error())
			return // 终止程序
		}
		go terry(conn)

		go tom(conn)
	}
}

var ch = make(chan int)

func terry(conn net.Conn) {
	for {
		buf := make([]byte, 512)
		len, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading", err.Error())
			return
		}
		s := string(buf[:len])
		l := strings.Split(s, ":")
		if l[0] == "terry" {
			fmt.Printf("receive: %s\n", buf[:len])
			conn.Write([]byte("你好, 我是terry"))
		} else {
			ch <- 1
		}

	}
}

// tom 老睡觉, ch 叫醒才能起来, 如果次数太多 会生气
func tom(conn net.Conn) {
	var total int
	for {
		select {
		case i := <-ch:
			total += i
			if total > 5 {
				conn.Write([]byte("tom: 我生气了!"))
				total = 0
			}
			conn.Write([]byte("tom: 叫我干嘛?"))
		}
	}
}
