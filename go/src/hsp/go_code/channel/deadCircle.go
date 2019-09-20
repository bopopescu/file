package main

import(
	"fmt"
)

func main(){
	// 创建3个管道
	var intchan chan int = make(chan int,10)
	close(intchan)

	// 让主线程一直循环
	go func(){
		for i:=0;i<4;i++{
			<-intchan
			fmt.Println("在读数据...")
		}
	}()
	


	fmt.Println("主程序结束了")
}