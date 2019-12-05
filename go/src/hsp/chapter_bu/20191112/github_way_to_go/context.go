package main

import (
	"fmt"
	"time"
	"context"
)
func main() {
	//test()



	//return
	ctx, cancel := context.WithCancel(context.Background())
	go watch(ctx,"【监控1】")
	go watch(ctx,"【监控2】")
	go watch(ctx,"【监控3】")
	time.Sleep(10 * time.Second)
	fmt.Println("可以了，通知监控停止")
	cancel()
	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(5 * time.Second)
}
func watch(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name,"监控退出，停止了...")
			return
			//todo 如果注释掉return  也会不停的打印   因为cancel相当于了closechan
		default:
			fmt.Println(name,"goroutine监控中...")
			//time.Sleep(2 * time.Second)
		}
	}
}

type Student struct {
	Name string
}

func test(){


	fmt.Printf("r2.leftUp 本身地址=%p \n", &Student{Name: "menglu"})

		fmt.Println(&Student{Name: "menglu"} == &Student{Name: "menglu"})
		fmt.Printf("%p,%p",&Student{Name: "menglu"},&Student{Name: "menglu"} )
	fmt.Println()
		fmt.Println(Student{Name: "menglu"} == Student{Name: "menglu"})

}