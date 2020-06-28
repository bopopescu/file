package main

import (
	"fmt"
	"time"
)
var ji chan int
func main() {
	ji =make(chan int)
	fmt.Println(123)


	go func() {

		time.Sleep(1e9)
		StopConsume()
	}()
	<-ji


	return
}
// StopConsume : 停止监听队列
func StopConsume() {
	ji <- 1
}