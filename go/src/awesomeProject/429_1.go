package main

import (
	"fmt"
)

func main(){
	//var a int
	//var b string
	const ask string="dsd"
	//var x,y,_ int
	myfunc()

	const (io1=iota;io2=iota)
	const (faker  int =0;theshy string = "khan")
	var s string="wangfeihu"
	var c = []rune(s)
	c[0]='c'
	//s2 := string(c)
	//var s2 string =c
	fmt.Println(string(c),"\n")

	fmt.Println(faker,theshy,io1,io2,"s'")

	if 2.01==2{
		fmt.Println("yes")

	}else
	{
		fmt.Println("nonono")
	}


	forfunc()
	//fmt.Println(b,123,a,x,y,ask,io1,io2)
	//fmt.Printf("sdas %s","dasdasda")

	fmt.Println(len("sadssssssss"))
}

func myfunc(){
	i:=0
	Here:
	fmt.Println(i)
	i++
	if(i==10){
		return
	}
	goto Here


}

//test  goland git

//reversation
func forfunc(){

	for i:= 1;i<5;i++{
		fmt.Println(i)
	}

	var s ="abcdefg"
	var a = []rune(s)

	for i,j:=0,len(a)-1;i<j;i,j=i+1,j-1{
		if(i==2){
			continue
		}
		a[i],a[j]=a[j],a[i]
	}

	fmt.Println(a)
	fmt.Println(string(a))
}