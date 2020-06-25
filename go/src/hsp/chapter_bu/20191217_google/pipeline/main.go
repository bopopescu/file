package pipeline

import (
	"encoding/binary"
	"fmt"
	"io"
	"math/rand"
	"sort"
	"sync"
	"time"
)

var starttime time.Time

func Init() {
	starttime = time.Now()
}

//读取数据
func ReaderSource(reader io.Reader, chunkSize int) <-chan int {
	out := make(chan int)
	//out := make(chan int,1024)  //缓冲 优化效率

	go func() {
		buffer := make([]byte, 8)
		bytesread := 0
		for {
			n, err := reader.Read(buffer) //n 是上面 那个 buffer 8 字节
			bytesread += n

			//fmt.Println(n,err)
			if n > 0 {
				v := int(binary.BigEndian.Uint64(buffer))
				out <- v
			}
			if err != nil || (chunkSize != -1 && bytesread >= chunkSize) {
				break
			}
		}
		close(out)
	}()

	return out
}

//写入文件
func WriteSink(writer io.Writer, in <-chan int) {

	time.Sleep(4e9)
	for v := range in {
		buffer := make([]byte, 8)
		binary.BigEndian.PutUint64(buffer, uint64(v))

		writer.Write(buffer)
		//fmt.Println("write down ->",uint64(v))
	}

}

func Merge(chan1, chan2 <-chan int) <-chan int {
	out := make(chan int)
	fmt.Println("chan1,chan2", chan1, chan2)
	//out := make(chan int,1024)  //缓冲 优化效率
	go func() {
		v1, ok1 := <-chan1
		v2, ok2 := <-chan2
		famous := ""

		fmt.Println(v1, ok1, v2, ok2)

		for ok1 || ok2 {
			//fmt.Println(ok1,ok2,v1,v2)
			//如果只有 v1  或者  v1 和 v2 都有 但是v1<v2 那么 v1入结果列
			//一般来讲  v1 和v2都是有的  如果只有1者  表示另一个chan已全部处理进入结果列
			if !ok2 || (ok1 && v1 <= v2) {
				out <- v1
				famous = fmt.Sprintf("v1  is %d", v1)
				v1, ok1 = <-chan1
			} else {
				out <- v2
				famous = fmt.Sprintf("v2  is %d", v2)

				v2, ok2 = <-chan2
			}
			_ = famous
			//log.Println("famous",famous)
		}

		close(out)

		fmt.Println("merge done", time.Now().Sub(starttime))

		fmt.Println("close")
	}()

	return out
}

func MergeN(inputs ...<-chan int) <-chan int {

	if len(inputs) == 1 {
		return inputs[0]
	}

	m := len(inputs) / 2

	return Merge(MergeN(inputs[:m]...), MergeN(inputs[m:]...))

}

//内部排序
func InMemSort(in <-chan int) <-chan int {

	out := make(chan int)
	//out := make(chan int,1024)  //缓冲 优化效率  800m数据  2min => 44s   80m  25s => 7s
	go func() {
		chantoarr := []int{}
		for v := range in {
			//fmt.Println("inMemsort console =>",v)
			chantoarr = append(chantoarr, v)
		}
		fmt.Println("read done", time.Now().Sub(starttime))

		sort.Ints(chantoarr)

		fmt.Println("sort done", time.Now().Sub(starttime))

		for _, v := range chantoarr {
			out <- v
		}
		close(out)
	}()

	return out
}

//生成随机数
func RandomSource(count int) <-chan int {
	out := make(chan int)

	go func() {
		for i := 0; i < count; i++ {
			out <- rand.Int()
		}
		close(out)
	}()

	return out
}

///////////////////////////////////////////////////////
///////////////////////////////////////////////////////
///////////////////////////////////////////////////////
////////////////////一些尝试和基本废弃的方法///////////////////////////////////

func ArraySource(a ...int) <-chan int {
	out := make(chan int)

	go func() {
		for _, v := range a {
			out <- v
			time.Sleep(1e9)
			//fmt.Println(666)
		}
		close(out)
	}()

	fmt.Println(321)
	return out
}

func RandomSourcego(count int) <-chan int {
	out := make(chan int, 200)

	go func() {
		for i := 0; i < count/2; i++ {
			out <- rand.Int()
		}
		close(out)
	}()
	go func() {
		for i := 0; i < (count/2)-100; i++ {
			out <- rand.Int()
		}
	}()

	return out
}

func WriteSinkgo(writer io.Writer, in <-chan int, wg *sync.WaitGroup) {

	go gowrite(writer, in, 1, wg)
	go gowrite(writer, in, 2, wg)

	/*	for i:=0;i<=1;i++{




		go func() {
			for v:=range in{
				buffer := make([]byte,8)
				binary.BigEndian.PutUint64(buffer,uint64(v))
				writer.Write(buffer)
				fmt.Println("go",i)
			}

			wg.Done()
		}()

	}*/

}

func gowrite(writer io.Writer, in <-chan int, no int, wg *sync.WaitGroup) {

	for {
		v, ok := <-in
		if ok == false {
			break
		}
		buffer := make([]byte, 8)
		binary.BigEndian.PutUint64(buffer, uint64(v))
		//writer.Write(buffer)

		fmt.Println("goroutine", no)

	}

	wg.Done()

}
