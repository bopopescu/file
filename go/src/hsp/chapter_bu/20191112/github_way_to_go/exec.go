package main

import (
	"fmt"
	"sync"
)

type query func(string) string

func exec(name string, vs ...query) string {
	ch := make(chan string, 10)
	wg := sync.WaitGroup{}
	wg.Add(4)
	fn := func(i int) {
		ch <- vs[i](name)
		wg.Done()
	}

	for i, _ := range vs {
		fmt.Println(i)
		go fn(i)

	}
	wg.Wait()
	close(ch)

	wg.Add(1)

	fmt.Println(2132)
	go func() {
		wg.Done()

		for o := range ch {

			fmt.Println(o)
		}

		fmt.Println("dsads")

	}()
	wg.Wait()

	//time.Sleep(1e9)
	return ""
	//return <-ch
}

func main() {
	ret := exec("111", func(n string) string {
		return n + "func1"
	}, func(n string) string {
		return n + "func2"
	}, func(n string) string {
		return n + "func3"
	}, func(n string) string {
		return n + "func4"
	})
	fmt.Println(ret)
}
