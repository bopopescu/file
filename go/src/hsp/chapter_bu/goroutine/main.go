package main

import (
	"fmt"
	"runtime"
	"strconv"
	"time"
)

func test() {
	for i := 0; i < 10; i++ {
		fmt.Println("go() -> hello"+strconv.Itoa(i), "world")
		time.Sleep(time.Second)
	}

}

func main() {

	cpus := runtime.NumCPU()
	fmt.Println(cpus)
	runtime.GOMAXPROCS(3)
	fmt.Println(123)
	return
	fmt.Println("hello"+string("asd"), "is thar goroutine  ")

	btime := time.Now().Unix()
	fmt.Println(btime, 666)

	go test()
	for i := 0; i < 10; i++ {
		fmt.Println("main () ->hello php")
		time.Sleep(time.Second)
	}
	etime := time.Now().Unix()
	fmt.Println(etime)

	fmt.Printf("%v", etime-btime)
}
