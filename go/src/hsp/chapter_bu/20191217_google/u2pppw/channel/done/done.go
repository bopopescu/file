package main

import (
	"fmt"
	"sync"
)

func doWork(id int,
	w worker) {
	for n := range w.in {
		fmt.Printf("Worker %d received %c\n",
			id, n)
		//w.done()
		go func() {
			fmt.Println("ffffffffffff") //todo 第一个就已经坏了哦

			w.done <- true
		fmt.Println("zzzzzzzzzzz") //todo 第一个就已经坏了哦

		}()
	}
}

type worker struct {
	in chan int

	//done func()
	done chan bool
}

func createWorker(
	id int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		//done: func() {
		//	wg.Done()
		//},
		done: make(chan bool),
	}
	//go doWork(id, w)

	go func() {
		for n := range w.in {
			fmt.Printf("Worker %d received %c\n",
				id, n)
			//w.done()
			go func() {
				//fmt.Println("ffffffffffff") //todo 第一个就已经坏了哦

				w.done <- true
				fmt.Println("zzzzzzzzzzz") //todo 第一个就已经坏了哦

			}()
		}
	}()
	return w
}

func chanDemo() {
	var wg sync.WaitGroup

	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}
	//wg.Add(20)
	for i, worker := range workers {
		worker.in <- 'a' + i
	}


	for _, worker := range workers {
		<-worker.done
	}
	for i, worker := range workers {
		worker.in <- 'A' + i
	}

	for _, worker := range workers {
		<-worker.done
	}
//todo 根本原因  不是协程就会 deadlock  （协程里面也要协程done）
//	for _, worker := range workers {
//		<-worker.done
//		<-worker.done
//	}

	//time.Sleep(time.Minute)
	return
	//wg.Wait()
}

func main() {
	chanDemo()
}
