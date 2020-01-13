package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {

	chan1 := make(chan int, 10)

	for i := 1; i <= 10; i++ {
		chan1 <- i
	}

	chan2 := make(chan string, 5)

	for i := 0; i < 5; i++ {
		chan2 <- "are u ok" + strconv.Itoa(i)
	}

	for {

		select {
		case v := <-chan1:
			time.Sleep(time.Second * 1 / 2)

			fmt.Println("632", v)

		case v := <-chan2:
			time.Sleep(time.Second * 1 / 2)

			fmt.Println("str ok ", v)
		default:
			fmt.Println("all done")
			return
		}

	}
}
