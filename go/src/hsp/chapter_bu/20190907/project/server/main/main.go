package main

import (
	"fmt"
	"net"
)

func main() {

	fmt.Println("mvc 服务器监听8889端口")
	listen, err := net.Listen("tcp", "127.0.0.1:8889")
	if err != nil {
		fmt.Println("listen wrong", err)

		return
	}

	for {
		//等待客户端连接

		conn, err := listen.Accept()

		if err != nil {
			fmt.Println("accept wrong", err)
		}

		//连接成功  启动协程保持通讯
		go processmr(conn)
	}

}

func processmr(conn net.Conn) {
	//循环读客户端发送的信息
	defer conn.Close()

	thepr := &Processor{
		Conn: conn,
	}
	err := thepr.processmr()
	if err != nil {
		fmt.Println("sever wrong", err)
		return
	}
}
