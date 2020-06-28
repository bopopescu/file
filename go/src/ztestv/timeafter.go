package main

import (
	"fmt"
	"time"

)

func main() {

	fmt.Println(6666)

	 time.AfterFunc(3*time.Second, func() {
		fmt.Print(888)
	})

	fmt.Println(123)
	time.Sleep(4e9)
}


func Abc(){
	fmt.Println(123)
}