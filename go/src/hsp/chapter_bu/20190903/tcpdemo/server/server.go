package main

import (
	"fmt"
	"io"
	"net"
)

func process(conn net.Conn) {

	defer conn.Close()

	for {
		buf := make([]byte, 1024)
		//fmt.Println("正在等待客户端发送信息 ip",conn.RemoteAddr())
		//fmt.Println("waiting")
		n, err := conn.Read(buf) //todo thisis block阻塞
		if err != nil {

			if err == io.EOF {
				fmt.Println("客户端已经退出", conn.RemoteAddr(), err) //todo   notice here  EOF
				return
			}
			fmt.Println("read wrong,", err)
			//return
		}

		fmt.Println("收到客户端的信息", conn.RemoteAddr(), string(buf[:n]))
	}

}

func main() {

	fmt.Println("start listening")

	listen, err := net.Listen("tcp", "127.0.0.1:52889")

	if err != nil {
		fmt.Println("listen err", err)
		return
	}

	fmt.Println(listen.Addr())
	defer listen.Close()

	for {
		//wait client conn
		fmt.Println("wait...")

		conn, err := listen.Accept()

		if err != nil {
			fmt.Println("accept wrong")
		} else {
			fmt.Println("accept succ", conn, conn.RemoteAddr().String(), conn.RemoteAddr().Network())

			go process(conn)

		}
	}

	fmt.Println("listen succ", listen)

}
