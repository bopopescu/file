package main

import "fmt"

func main(){

	var a map[int]string
	a = make(map[int]string,10)
	a[1]="asd"
	a[10]="asd"
	a[110]="asd"
fmt.Println(a)


	b:= make([][]int,10)
	for i,v:=range b{
		fmt.Println(i,v)
	}
	fmt.Println(b)

}
