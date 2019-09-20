package main

import (
	"reflect"
	"fmt"
)
/**
todo 关键点  elem（）
 */


func main(){
	//通过反射修改值
	arr:=[]int{1,2,3}
	fmt.Println(arr)
	fmt.Printf("%T\n",arr)
	num:= 133
	reflectModify(&num)  //地址 &  ptr =》 指针
	fmt.Println(num)



	//类似的理解
	num3:=123
	var num4 *int = &num3
	fmt.Println(*num4)
}

func reflectModify(b interface{}){


	rVal:=reflect.ValueOf(b)

	fmt.Println(rVal.Kind())
	fmt.Println(rVal.Type())


	//rVal是个ptr指针   要转成值类型   理解成  取值  类似于  *取值
	rVal.Elem().SetInt(155)  //set必须use addressable 地址的
	//rVal.SetInt(155)   //todo wrong






}