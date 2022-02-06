package main

import (
	"fmt"
	"net"
	protof "app/proto"
)

// 黏包 client

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("dial failed, err", err)
		return
	}
	defer conn.Close()
	for i := 0; i < 20; i++ {
		msg := `Hello. How are you?`
		// 调用协议编码数据
		b, err := protof.Encode(msg)
		if err != nil {
			fmt.Println("encode failed,err:", err)
			return
		}
		conn.Write(b)
		// time.Sleep(time.Second)
	}
}
