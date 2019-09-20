package main

import(
	"fmt"
)

func writeData(intchan chan int){
	for i:=0;i<50;i++{
		intchan <- i
		fmt.Println("写入数据",i)
	}
	close(intchan)
}
func readData(intchan chan int,extchan chan bool){
	for {
		v,ok := <- intchan
		if ok!=true{
			break
		}
		fmt.Println("读取数据",v)
	}
	extchan <- true
	close(extchan)
}

func main(){
	// 创建2个管道
	var intchan = make(chan int,10)
	var extchan = make(chan bool,1)

	// 管道一直在写,没有读的话,就会阻塞
	go writeData(intchan)
	go readData(intchan,extchan)
	
	// 如果没有下面这个死循环,主程序一下子就结束了,协程也就结束了.
	for{
		_,ok := <- extchan
		if ok {
			break
		}
	}
}