package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

// tcp server端
func process(conn net.Conn) {
	defer conn.Close()
	// 3.与客户端通信
	reader := bufio.NewReader(conn)
	var buf [1024]byte
	for {
		n, err := reader.Read(buf[:])
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("read from client failed, err:", err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("收到client发来的数据:", recvStr)
	}

}

func main() {
	// 1.本地端口启动服务
	listen, err := net.Listen("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("start server on 127.0.0.1:30000 failed, err:", err)
		return
	}
	defer listen.Close()
	// 2.等待别人来跟我建立连接
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		go process(conn)
	}
}
