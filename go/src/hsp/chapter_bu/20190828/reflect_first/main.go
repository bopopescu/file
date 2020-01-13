package main

import (
	"reflect"

	"fmt"
)

func reflectTest(b interface{}) {

	//testv := b.(string)
	//testv := b.(int)
	//fmt.Println(testv)

	//first ref.type

	rtype := reflect.TypeOf(b)

	fmt.Println(rtype)

	rval := reflect.ValueOf(b)

	rkind := rval.Kind()
	fmt.Println(rkind, 123)

	fmt.Println(rval)

	//n2:=2+rval  //wrong
	n2 := 2 + rval.Int() //wrong todo thisis 取值

	fmt.Println(n2)
	fmt.Printf("%T\n", rval)

	iv := rval.Interface()

	fmt.Printf("%T\n", iv)

	n3 := 10 + iv.(int) //todo thisis 断言
	//n5:=10+iv
	num2 := iv.(int)
	fmt.Println(num2, n3)

}

func main() {

	var num int = 232

	reflectTest(num)

}
