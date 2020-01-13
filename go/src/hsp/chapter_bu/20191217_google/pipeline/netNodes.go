package pipeline

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"time"
)

func NetworkSink(addr string, in <-chan int) {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	go func() {
		defer listener.Close()
		conn, err := listener.Accept()

		if err != nil {
			panic(fmt.Sprintf("accept err %s", err))
		}
		defer conn.Close()

		//WriteSink(conn,in)

		writer := bufio.NewWriter(conn)

		defer writer.Flush()
		WriteSink(writer, in)

	}()

}

func NetworkSource(addr string) <-chan int {

	out := make(chan int, 0)

	go func() {
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			panic(err)
		}
		defer conn.Close()

		reader := bufio.NewReader(conn)

		r := ReaderSource(reader, -1)

		for v := range r {
			out <- v
		}
		close(out)
	}()
	return out

	//todo 直接这样竟然不行   find why！！！！！

	// todo 为什么会直接 merge done  而没要预期的阻塞？？？？？？？？？？？？？？？？？？？

	//TODO 然而上面的可以 ？？？？

	//todo  ！！！！基本确定了   应该是 defer close 的问题！！！！！！

	/*	conn, err := net.Dial("tcp", addr)
		if err != nil {
			panic(err)
		}
		//defer conn.Close()

		reader := bufio.NewReader(conn)


		out :=ReaderSource(reader,-1)   //why dont use this, or can only this work

		fmt.Println(out)
		return out
	*/

	//todo  下面这段代码回合又不行呢

	//todo  显然不行  gofunc 已经废了。。。   无限阻塞 空chan
	/*	out := make(<-chan int)
		go func() {
			conn, err := net.Dial("tcp",addr)
			if err != nil {
				panic(err)
			}

			fmt.Println(11111111111)
			// 原代码为何需要先得到 r ，再把r中的每一个数据传给 out， 为什么不能直接用out！
			out = ReaderSource(bufio.NewReader(conn),-1)
			fmt.Println(22222222,out)
			}()

		fmt.Println("return ,",out)
		return out*/
}

func CreateNetPipeline(filename string, fileSize, chunkCount int) <-chan int {

	chunkSize := fileSize / chunkCount

	Init()
	//sourceRes := make([]<-chan int,chunkCount)
	sortAddr := []string{}
	log.Println("----------begin create", time.Now().Unix())
	for i := 0; i < chunkCount; i++ {
		file, err := os.Open(filename)
		if err != nil {
			panic(err)
		}

		file.Seek(int64(i*chunkSize), 0)

		source := ReaderSource(bufio.NewReader(file), chunkSize)

		//sourceRes = append(sourceRes,InMemSort(source))

		addr := ":" + strconv.Itoa(7000+i)

		//相当于go协程直接开了
		NetworkSink(addr, InMemSort(source))

		sortAddr = append(sortAddr, addr)

	}

	//return nil

	//相当于 client 端     上面的是服务端
	sourceRes := []<-chan int{}
	for i, addr := range sortAddr {
		sourceRes = append(sourceRes, NetworkSource(addr))

		fmt.Println("the addr index is ", i)

	}

	log.Println("----------end createNet", time.Now().Unix())
	fmt.Println("source 长度是", len(sourceRes))
	return MergeN(sourceRes...)
}
