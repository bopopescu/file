package main

import(
	"fmt"
	"strconv"
	"time"
    "sync"
)

var (
	myMap = make(map[int]int,10)
	// 声明一个全局的互斥锁
	lock sync.Mutex

)

func startRoutine(){
	for i:=0;i<10;i++{
		fmt.Println("test()"+strconv.Itoa(i))
		time.Sleep(time.Second);
	}
}

func syncMethod(n int){
	res :=1
	for i:=1;i<=n;i++{
		res *= i
	}
	// 加锁
	lock.Lock()
	myMap[n]=res;
	lock.Unlock()
}

func main(){

	// 在主线程上开启一个协程. 如果主线程退出了,协程即使没有执行完也会退出.
	// goroutine的调度模式:MPG模式,M主线程,P上下文,G协程
	// 假设有3个主线程都运行在一个cpu上,叫并发,
	// 假设3个主线程运行在不同的cpu上,就叫并行
	

	// 案例1:开启协程
	//go startRoutine();



	// 案例2:
	// go build -race demo.go 查看资源竞争关系
	for i:=1;i<=20;i++{
		go syncMethod(i);
	}
	time.Sleep(time.Second*5); // 这能睡眠的时间你根本不能决定?从而引出管道的概念.

	lock.Lock()
	for k,v := range myMap{
		fmt.Printf("%v的阶层是%v \n",k,v)
	}
	lock.Unlock()


}