package main

import (
	"bufio"
	"fmt"
	"hsp/chapter_bu/20191217_google/pipeline"
	"math/rand"
	"os"
	"sync"
	"time"
)

//目前是生成文件初始   externalSort/main 是直接拿这个生成的large.in去跑
func main() {

	const filename = "large.in"
	const n = 10000000

	timestart := time.Now()
	timestart2 := time.Now().Unix()

	const usego = 0

	file, e := writeWorker(filename, n, usego)

	timeend := time.Now().Unix()

	duration := time.Now().Sub(timestart)
	readWorker(file, e, filename)

	fmt.Println(file, e, duration, timeend-timestart2)

	return

	closetest()

	return

	testclose()

	Mergetest()

	Range2For()
}
func closetest() {
	ch1 := make(chan int)
	ch2 := make(chan bool)

	go write(ch1)
	go read(ch1, ch2)

	<-ch2

}

func read(ch1 chan int, ch2 chan bool) {
	for {
		v, ok := <-ch1
		if ok == false {
			ch2 <- true
		} else {
			fmt.Printf("read a int is %d\n", v)

		}

	}
}

func write(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

func readWorker(file *os.File, e error, filename string) {
	file, e = os.Open(filename)
	if e != nil {
		panic(e)
	}
	defer file.Close()
	fmt.Println(file, e)
	readchan := pipeline.ReaderSource(bufio.NewReader(file), -1)
	fmt.Println("testme is ", rand.Int31())
	aftersort := pipeline.InMemSort(readchan)
	count := 0
	for v := range aftersort {
		count++
		fmt.Println(v)
		if count >= 100 {
			break
		}
	}
}

func writeWorker(filename string, n int, gowrite int) (*os.File, error) {
	file, e := os.Create(filename)
	if e != nil {
		panic(e)
	}
	defer file.Close()
	//file is writer

	if gowrite == 1 {
		//协程同时写？
		writeChan := pipeline.RandomSourcego(n)

		wg := sync.WaitGroup{}
		wg.Add(2)

		pipeline.WriteSinkgo(file, writeChan, &wg)
		wg.Wait()

	} else {
		writeChan := pipeline.RandomSource(n)

		write := bufio.NewWriter(file)
		pipeline.WriteSink(write, writeChan)
		write.Flush() //用了bufio 缓冲最后得加flush
	}
	fmt.Println(file, e)
	return file, e
}

func Range2For() {
	/*for {
		if num,ok:=<-ch;ok{
			fmt.Println(num)
		}else{
			break
		}
	}*/
}

func testclose() {
	/*
		testchan := make(chan int)

		go func() {

			v1,ok1 :=<- testchan
			fmt.Println(v1,ok1)
			v2,ok2 :=<- testchan
			fmt.Println(v2,ok2,66)


			//fmt.Println(123)
			v3,ok3 :=<- testchan
			fmt.Println(v3,ok3)
			fmt.Println(1234)

		}()

		testchan<-1

		testchan<-2

		close(testchan)   //不close 携程内会阻塞

		time.Sleep(time.Microsecond*10000)
		time.Sleep(1e9)



		return
	*/
	//ch :=pipeline.InMemSort(
	//	pipeline.ArraySource(1,3,123,35,11,23,120000,3144))
}

func Mergetest() {
	ch := pipeline.Merge(
		pipeline.InMemSort(
			pipeline.ArraySource(100, 13, 15, 17, 123, 321)),
		pipeline.InMemSort(
			pipeline.ArraySource(2, 4, 6, 8)))
	fmt.Println(ch)
	for v := range ch {
		fmt.Println("paixuhou-------", v)
	}
}
