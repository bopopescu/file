package main

import (
	"fmt"
	"math"
)

var a,b int = 3,4

func triangle(){
	var c int
	c=int(math.Sqrt(float64(a*a+b*b)))

	fmt.Println(c)
}


var aa int =3
var ss ="kk"

var (a1=3
a2=33
a3=333
)
//dasdasd
func variable(){
	var first int
	var second string
	fmt.Printf("%d %q\r\n",first,second)
}

func varcvalue(){
	var a,b int =3,4
	var s string ="asd"
	fmt.Println(a,b,s)
}
func main(){
	fmt.Println("dddd")
	variable()
	fmt.Println("sdddd")
	varcvalue()
	test()
	triangle()
}

func test(){
	var a,b,c,d=1,2,3,"ds"
	b=1221
	fmt.Println(a,b,c,d,ss,aa+1,a3,math.Pi)
}