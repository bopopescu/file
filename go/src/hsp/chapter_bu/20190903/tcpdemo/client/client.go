package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:52889")

	if err != nil {
		fmt.Println("dial err", err)
		return
	}

	fmt.Println("dial succ", conn)

	defer conn.Close()


	reader := bufio.NewReader(os.Stdin)

	for {

		content, err := reader.ReadString('\n')

		if err != nil {

			fmt.Println("getcontent wrong,", err)
		}

		fmt.Println("your type is ", content)

		content = strings.Trim(content, " \r\n")
		if string(content) == "exit" {
			//conn.Close()
			return
		}

		contetns := "exit"
		fmt.Printf("%T %v %v %v\n ", content, content == "exit", contetns == "exit", content == contetns)

		n, err := conn.Write([]byte(content + "\n"))

		if err != nil {

			fmt.Println("write wrong,", err)
		}

		fmt.Println("send byte  num", n)
	}

}
