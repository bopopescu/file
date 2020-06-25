package main

import (

"fmt"
"log"
"sort"
"sync"

)

var slice []int
var slice2 []int

var Logger *log.Logger
func main() {

	logArr := make([]int, 0)

	fmt.Println(logArr)
	//return

	wg:=sync.WaitGroup{}
	//wg.Add(100)
	slice2= make([]int,200)

	mlk:=sync.Mutex{}
	for i:=0;i<100;i++{
		wg.Add(1)//add 的肯定有100个  但是操作共享变量会导致线程不安全的情况

		//fmt.Println(i)
		slice2[i]=i
		go func(i2 int) {
			//wg.Add(1)//add 的肯定有100个  但是操作共享变量会导致线程不安全的情况

			fmt.Println(i2,"====")
			defer wg.Done()
			mlk.Lock()
			defer mlk.Unlock()

			slice=append(slice,i2)

		}(i)



	}


	wg.Wait()

	p:=0
	for i,v:=range slice2{
		if v>0 || i==0{
			fmt.Println(i,v)

	p++
		}
	}
	fmt.Println(p)
	//fmt.Println(len(slice2))
	//time.Sleep(1e9)
	fmt.Println()

	fmt.Println(len(slice))
	fmt.Println()

	sort.Ints(slice)
	fmt.Print(slice)
}
