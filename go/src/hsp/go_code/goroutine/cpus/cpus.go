package main

import (
	"fmt"
	"runtime"
)

func main() {
	cpuNum := runtime.NumCPU()
	fmt.Println("cpu个数是", cpuNum)

	//自己可是设置使用cpu的个数,
	runtime.GOMAXPROCS(cpuNum - 1)
}
