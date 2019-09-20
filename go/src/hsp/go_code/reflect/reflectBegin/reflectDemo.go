package main

import(
	"fmt"
	"reflect"
)


func reflectTest01(i interface{}){
	// 通过反射获取传入的变量的type
	// 1,先获取reflect.Type
	rtype := reflect.TypeOf(i)
	fmt.Println("rtype的类型是",rtype)

	// 2,获取reflect.Value
	rVal := reflect.ValueOf(i)
	fmt.Println("rVal是",rVal)

	// 3,将rVal转成interface{}
	iv := rVal.Interface()
	// 4,将interface{}转成需要的类型:断言
	num := iv.(int)
	fmt.Println("num是",num)

}

func main(){
	// 变量,interface{},value之间可以互相转化
	var num int =100
	reflectTest01(num)


}