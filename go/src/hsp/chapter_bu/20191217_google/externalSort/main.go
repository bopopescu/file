package main

import (
	"bufio"
	"fmt"
	"hsp/chapter_bu/20191217_google/pipeline"
	"log"
	"math/rand"
	"os"
	"time"
)

func fuck() <-chan int {
	out := make(chan int)

	go func() {
		out <- 333
		time.Sleep(2e9)
		close(out)
	}()

	return out
}

func you(chan1 <-chan int) <-chan int {
	out := make(chan int)
	//out := make(chan int,1024)  //缓冲 优化效率
	go func() {
		v1, ok1 := <-chan1

		fmt.Println(v1, ok1)

		out <- 1321

		close(out)

		fmt.Println("close")
	}()

	return out
}

func main() {
	//testTodo()
	//_ = pipeline.CreateNetPipeline("large.in",80000000,4)
	//
	//time.Sleep(time.Hour)
	//return

	timenow := time.Now()

	//虽然这个页面都没有写 go   但是实际上全是协程同步在跑   chan

	//单机版外部排序
	//p:=createPipeline("small.in",5120000,4)

	//网络版外部排序
	p := pipeline.CreateNetPipeline("large.in", 80000000, 4)

	//80000000 = 10000000 * 8
	//p:=createPipeline("large.in",80000000,4)
	//_=p
	//return

	//将排序好的写入文件 out1
	writeOutPipeline("large.out1", p)

	//打印这个文件   实际上在上面writeout打印也一样
	PrintFile("large.out1")

	fmt.Println(time.Now().Sub(timenow))

}

func testTodo() {
	testchan2 := fuck()
	p2 := you(testchan2)
	for v := range p2 {
		fmt.Println(v)
	}

	return
	testchan := make(chan int)
	fmt.Println(123)
	go func() {
		time.Sleep(1e9)
		close(testchan)
	}()
	v1, ok := <-testchan
	fmt.Println(v1, ok)
}
func PrintFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	readchan := pipeline.ReaderSource(bufio.NewReader(file), -1)
	fmt.Println("testme is ", rand.Int31())
	ct := 0
	for v := range readchan {
		ct++
		if ct > 100 {
			break
		}
		fmt.Println("askuy", v)
	}
}

func writeOutPipeline(filename string, p <-chan int) {

	file, e := os.Create(filename)
	if e != nil {
		panic(e)
	}
	defer file.Close()

	write := bufio.NewWriter(file)
	defer write.Flush()

	pipeline.WriteSink(write, p)
}

func createPipeline(filename string, fileSize, chunkCount int) <-chan int {

	chunkSize := fileSize / chunkCount

	pipeline.Init()
	//sourceRes := make([]<-chan int,chunkCount)
	sourceRes := []<-chan int{}
	log.Println("----------begin create", time.Now().Unix())
	for i := 0; i < chunkCount; i++ {
		file, err := os.Open(filename)
		if err != nil {
			panic(err)
		}

		file.Seek(int64(i*chunkSize), 0)

		source := pipeline.ReaderSource(bufio.NewReader(file), chunkSize)
		/**/
		fmt.Println("jjjjj", i, time.Now().Unix())
		sourceRes = append(sourceRes, pipeline.InMemSort(source))
		//inMsource :=pipeline.InMemSort(source)

	}

	log.Println("----------end create", time.Now().Unix())
	fmt.Println("source 长度是", len(sourceRes))
	return pipeline.MergeN(sourceRes...)
}
