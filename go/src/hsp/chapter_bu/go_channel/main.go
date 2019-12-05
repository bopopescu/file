 package main

import (

	"fmt"

	//"time"
	"sync"
	"time"
)

var myMap = make(map[int]int,10)
//
var lock sync.Mutex  //todo 加一把锁 全局互斥

func calJc(n int)int{

	res:=1
	for i:=1;i<=n;i++{
		res*=i
	}
	lock.Lock()
	myMap[n]=res
	lock.Unlock()
	return res
}
//计算1-200的阶乘   放入map中
func main(){

	//200 个协程
	for i:=1;i<=200;i++{
		go calJc(i)
	}

	//todo 如果不休眠  则主线程退出  协程也退出
	time.Sleep(time.Second *10 )

	// todo 如果不休眠  则主线程退出  协程也退出



	//todo为了避免资源竞争
	lock.Lock()
	for i,v:=range myMap{

		fmt.Printf("mao[%d]=%d\n",i ,v)
	}
	lock.Unlock()
}