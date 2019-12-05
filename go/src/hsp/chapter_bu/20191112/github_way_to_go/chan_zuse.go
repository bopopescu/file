package main

import (
	"fmt"
	"time"
)

func main() {




	test := make(chan int,3)
	test<-1
	test<-13


	//close(test)  //不关闭的话就会阻塞哦
	//
	// todo  如果 不在协程里  for range  底层就会判断 爆出deadclock
	fmt.Println("asdd")


	go func() {
		for i:=range(test){
			fmt.Println(i)

		}

		fmt.Println("协程里   出现这个代表 没有阻塞  就是close了")
	}()
	time.Sleep(2e9)

}