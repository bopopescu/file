package main

import(
	"fmt"
)

func closeChannel(){
	intchan := make(chan int,3)
	intchan <- 10
	intchan <- 20
	close(intchan)
	// 不能再写了,只能读
	num := <- intchan
	fmt.Println("",num)
}

func traverseChannel(){
	intchan := make(chan int,100)
	for i:=0;i<100;i++{
		intchan <- i*2
	}
	// 遍历时要先关闭,否则会报死锁的错误
	close(intchan)
	for v := range intchan{
		fmt.Print("\t",v)
	}
	
}

func traverseChanne2(){
	intchan := make(chan int,10)
	intchan1 := make(chan int,10)
	for i:=0;i<10;i++{
		intchan <- i*2
	}
	// 遍历时要先关闭,否则会报死锁的错误
	// 使用select,就可以不用关闭管道了
	for{
		select{
		case v := <- intchan :
			fmt.Print("\t",v) 
		case v := <- intchan1 :
			fmt.Print("\t",v)
		default:
			fmt.Print("都取不到数据啦")
			return
			// break
		}
	}
	
}
func main(){
	// closeChannel()

	// 遍历
	//traverseChannel()

	// 遍历
	traverseChanne2()
	

	
}