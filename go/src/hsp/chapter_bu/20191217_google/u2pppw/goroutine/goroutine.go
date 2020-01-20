package main

import (
	"fmt"
	"time"
)

func main() {

	var a [10]int
	for i := 0; i < 10; i++ {

		go func(i int) {
			for {
				a[i]++
				//runtime.Gosched() //TODO 手动教出控制权
			}
		}(i)
		//
		//go func() {
		//	for{
		//		a[i]++//TODO 这个bug是因为 i都是10   因为是 ！！！闭包！！！
		//		//TODO 闭包 都绑定了一个变量
		//		runtime.Gosched()
		//	}
		//}()
	}
	fmt.Println(123)

	time.Sleep(time.Millisecond) //TODO  为何会死循环  跳不出  交不出控制权？

	fmt.Println(123321)
	fmt.Println(a)
	//for i := 0; i < 1000; i++ {
	//	go func(i2 int) {
	//		for {
	//			fmt.Printf("Hello from "+
	//				"goroutine %d\n", i2) //todo  这里的print 算是交出了控制权
	//		}
	//	}(i)
	//}
	//time.Sleep(time.Minute)
}
