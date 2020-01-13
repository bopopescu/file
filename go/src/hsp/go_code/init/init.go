package main

import (
	"fmt"
)

var fun = test()

// init函数用来完成初始化工作
func test() int {
	fmt.Println("全局变量 被执行了")
	return 0
}
func init() {
	fmt.Println("init函数 被执行了")
}
func main() {
	fmt.Println("main函数 被执行了")
}
