package main

import (
	"fmt"
	_"time"

	"time"
)

func MyRoutineFunc(ch chan int) {
	// 函数处理
	ch <- 123
	time.Sleep(time.Second*1)

	fmt.Println("MyRoutineFunc process finished.")
}

func test(){
	fmt.Println("123")
}
func main() {

	//var mychan = make(chan int,12)
	var mychan = make([]chan int, 10)

	mychan[1] = make(chan int,2)  //注意这个1   不要协程得话必须写  有协程竟然不用写
	//go MyRoutineFunc(mychan[1])   //协程这里面竟然不用指定chan的长度就可以 插进去
	 MyRoutineFunc(mychan[1])
	//mychan[1]<-123
	//close(mychan[1])
	rb,ok:= <-mychan[1]  //阻塞  协程里sleep了  这里也等了1秒  但是如果协程去掉的话  会直接报错  说明底层是知道  到底有没程序会操作这个chan的
	//<-mychan[1]
	//rb2,ok:= <-mychan
	fmt.Println(rb,ok,66666666,
	//,rb2
	)
return

	chs := make([]chan int,10)

	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go MyRoutineFunc(chs[i])  //? 为什么这个打印不出来呢
		//go test()

		//time.Sleep(time.Second * 1/10)

	}

	for k, ch := range chs {

		i,ok:=
			<-ch
		//_,ok2:=
		//	<-ch
		fmt.Println(i,"--",k,ok)
	}

	fmt.Println("All goroutine finished.")
}