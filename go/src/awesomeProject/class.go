package main

import (
	"fmt"
)

type nba struct {
	a string
	b int
}

func (f *nba) test(){
	fmt.Println(f.a)
	fmt.Println(f.b)
}




type Test struct {
	 a string
}

func (t Test) show(){
	fmt.Println(t.a)
}

func newTest()(test Test){
	test=Test{a:"durant"}
	test.a = "dssddasdasd"

	return  test
}


func main(){

b:=nba{a:"james harden",b:2222}
b.test()

c:=newTest()
c.show()
}
