package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	//打开连接:
	conn, err := net.Dial("tcp", "localhost:50000")
	if err != nil {
		//由于目标计算机积极拒绝而无法创建连接
		fmt.Println("Error dialing", err.Error())
		return // 终止程序
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("你想叫谁说话? Tom 还是 Terry? 格式 'Terry:{内容}'")
		input, _ := reader.ReadString('\n')
		trimmedInput := strings.Trim(input, "\r\n")
		if trimmedInput == "Q" {
			return
		}
		conn.Write([]byte(trimmedInput)) //测试忽略错误

		//开始读取内容
		buf := make([]byte, 512)
		len, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading", err.Error())
			return
		}

		fmt.Fprintf(os.Stdout, "receive:"+string(buf[:len])+"\r\n")
	}
}
