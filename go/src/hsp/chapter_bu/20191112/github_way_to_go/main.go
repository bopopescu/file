package main

import (
	"fmt"
	"strconv"
	"regexp"
	"sync"
	"time"
)
//the way to go小练习——闭包实现斐波那契数列
func fib() func() int {
	var a int = 1
	var b int = 1
	return func() int {
		c := a
		a = b
		b = a + c
		return c
	}
}
var xx  = 0
func increment3( ch chan bool) {
	ch <- true
	xx = xx + 1
	<- ch
}
func test3() {
	ch := make(chan bool, 1)
	for i := 0; i < 1000; i++ {
		fmt.Println(i)
		go increment3(ch)
	}

	//time.Sleep(time.Second)
	fmt.Println("final value of x", xx)
}

func main() {
	test3()
	return
	ch := make(chan int, 3)
	for i:=0;i<30;i++{
		go func(){
			fmt.Println(122)
			ch<-i
		}( )

	}

	time.Sleep(time.Second/10)

	//ch<-1
	//ch<-2
	//close(ch)
	fmt.Printf("intchan管道的值是%v,intchan管道的地址是%p \n",ch,&ch)

	fmt.Printf("intchan管道的长度是%v,intchan管道的容量是%v \n",len(ch),cap(ch))

	fmt.Println((ch),len(ch))
	return
	test2()
	f := fib()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}


func regx(){
		//目标字符串
		searchIn := "John: 2578.34 William: 4567.23 Steve: 5632.18"
		pat := "[0-9]+.[0-9]+" //正则

		f := func(s string) string{
			v, _ := strconv.ParseFloat(s, 32)
			return strconv.FormatFloat(v * 2, 'f', 2, 32)
		}

		if ok, _ := regexp.Match(pat, []byte(searchIn)); ok {
			fmt.Println("Match Found!")
		}

		re, _ := regexp.Compile(pat)
		//将匹配到的部分替换为"##.#"
		str := re.ReplaceAllString(searchIn, "##.#")
		fmt.Println(str)
		//参数为函数时
		str2 := re.ReplaceAllStringFunc(searchIn, f)
		fmt.Println(str2)
}

var x  = 0
func increment(wg *sync.WaitGroup, ch chan bool) {
	ch <- true
	x = x + 1
	<- ch
	wg.Done()
}
func test2() {
	var w sync.WaitGroup
	ch := make(chan bool, 1)
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment(&w, ch)
	}
	w.Wait()
	fmt.Println("final value of x", x)
}