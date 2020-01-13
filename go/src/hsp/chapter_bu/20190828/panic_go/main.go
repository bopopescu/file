package main

import (
	"fmt"
	"time"
)

func say() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second * 1 / 2)
		fmt.Println("are leijun ok")
	}
}

func test() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(" goroutine  test() wrong")
		}
	}()

	var mymap map[int]string
	//mymap = make(map[int]string,10)
	mymap[1] = "asd"

}

func main() {

	go say()
	go test()

	for i := 0; i < 10; i++ {
		time.Sleep(time.Second * 1 / 4)
		fmt.Println("are ttt ok")
	}
}
