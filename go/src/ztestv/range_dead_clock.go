package main

import (
	"fmt"
	"time"
)

func main() {

	test := make(chan int,3)
	test<-1
	test<-13
	test<-1323
	//test<-1333


	//close(test)  //不关闭的话就会阻塞哦   底层就会判断 爆出deadclock
	fmt.Println("asdd")

	//go func() {
	//	for{
	//		v,ok:=<-test
	//		if !ok{
	//			break
	//		}
	//		fmt.Println(v,ok)
	//	}
	//}()


	//
	for i:=range(test){
		fmt.Println(i,55555)
	}

	time.Sleep(1e9)
	fmt.Println("end")
}
