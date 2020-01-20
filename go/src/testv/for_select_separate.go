package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("begin")
	select {
	case<-time.NewTimer(1000*time.Millisecond).C:
		fmt.Println(123)
	//default:
	//	fmt.Println("ds")
	}


	//因为你在select外面加了一个for, 没有default时它会阻塞，直到enqueue所在的case执行， 加上default后，它会随机执行default和case，当case能正常执行时，随机执行到case会得出当前的结果。去掉select外的for可以实现你想要的结果	for{

	for{

		fmt.Println(232)
		select {
		case<-time.NewTimer(1000*time.Millisecond).C:
			//
			//default:
			//	fmt.Println(666)
		}
	}

}
