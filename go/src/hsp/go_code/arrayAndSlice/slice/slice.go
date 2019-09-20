package main

import (
	"fmt"
	"sort"
)
func slideCreat1(){
	var i [5]int = [5]int{2,15,55,25,35}
	sli := i[1:3]
	fmt.Println("slice元素是",sli)
	fmt.Println("slice元素的个数是",len(sli))
	fmt.Println("slice元素的容量是",cap(sli))
	
	// slice的地址
	fmt.Println("slice第一个元素的地址是",&sli[0])
}
func slideCreat2(){
	// 切片必须make后才能使用, 
	var slice []int = make([]int,5,10)
	// slice := make([]int,5,10)
	slice[0]=15
	fmt.Println("slice元素是",slice)
}
func slideCreat3(){
	// 切片必须make后才能使用, 
	var slice []int = []int{1,3,67}
	// slice := []int{1,3,67}
	slice[0]=15
	fmt.Println("slice元素是",slice)
}
func slideIterate1(){
	 
	var slice []int = []int{1,3,67}
	for i:=0;i<len(slice);i++ {
		fmt.Println("slice元素是",slice[i])
	}
}
func slideIterate2(){
	 
	var slice []int = []int{1,3,67}
	for _,v := range slice{
		fmt.Println("slice元素是",v)
	}
}
func slide2slice(){
	var i [5]int = [5]int{2,15,55,25,35}
	sli := i[1:4]
	sli2sli := sli[1:2]
	fmt.Println("sli2sli元素是",sli2sli)
}
func sliceAppend(){
	slice := []int{1,3,67}
	fmt.Println("slice元素是",slice)
	sl := append(slice,400,500)
	fmt.Println("sl追加后的元素是",sl)
}
func sliceCoppy(){
	slice1 := []int{1,3,67}
	slice2 := make([]int,6)
	copy(slice2,slice1)
	fmt.Println("slice2拷贝后的元素是",slice2)
}
func main(){
	// 切片最明显不同于数组的地方是:切片没有长度.它是引用类型,切片底层的数组是不可见的.所以改变slice也会改变数组的值
	// 切片底层是一个结构体(首地址,长度,容量),有3部分组成
	// slideCreat1()
	// slideCreat2()
	//slideCreat3()

	//切片的遍历和数组一样,也有2种方式
	// slideIterate1()
	// slideIterate2()

	//切片还可以继续切片
	// slide2slice()

	//切片动态追加
	sliceAppend()

	//切片拷贝,数组不能拷贝到切片. 拷贝的源与目标数据空间是独立的互不影响.
	sliceCoppy()

	//切片作为参数传递时,就像指针一样,会影响外面的值.

	//排序
	ss := make([]int,2)
	ss = append(ss,55)
	ss = append(ss,15)
	ss = append(ss,5)
	sort.Ints(ss)
	fmt.Println(ss)

}