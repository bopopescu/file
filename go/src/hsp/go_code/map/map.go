package main

import (
	"fmt"
	
)

func mapMake1(){
		// 光声明还不行,还要用make分配空间
		var a map[string]string
		a = make(map[string]string,10)
		a["name"]="xiangshousheng"
		fmt.Println("",a)
}
func mapMake2(){
	a := make(map[string]string)
	a["name"]="xiangshousheng"
	a["name1"]="ss"
	a["name2"]="dd"
	fmt.Println("",a)
}
func mapMake3(){
	a := map[string]string{
		"name":"xss",
		"sex":"male",
		"age":"30",
	}
	a["location"]="hubei"
	fmt.Println("",a)
}
func mapBestDemo(){
	// 请问你晕不晕?🐸😄
	a := make(map[string]map[string]string)
	a["location"]=make(map[string]string)
	if a["location"] != nil{
		fmt.Println("",a)
	}
	
}
func mapCrud(){
	a := make(map[string]string)
	// 增加,修改
	a["name"] = "xss"
	a["name"] = "xss1"
	a["age"] = "35"
	fmt.Println("",a)

	//使用go的内置函数来删除,如果要删除所有的key,直接给他make一个新的空间就行了
	// 即使没有这个key,你删除也不会报错
	delete(a,"name")
	
	
	fmt.Println("",a)

	//查找
	val,ok := a["age"]
	fmt.Println("val的值是",val,"ok的值是",ok)
}
func mapIterate(){
	a := make(map[string]string)
	a["name"] = "xss"
	a["age"] = "35"
	for k,v := range a {
		fmt.Println("k的值是",k,"v的值是",v)
	}

}
func mapSlice(){
	// 先声明一个切片
	var ghosts []map[string]string
	// 再分配空间
	ghosts = make([]map[string]string,2)

	ghosts[0] = make(map[string]string,3)
	ghosts[0]["age"] = "35"
	ghosts[0]["name"] ="xss"
	ghosts[0]["sex"] = "male"
	fmt.Println("值是",ghosts)
} 
func main(){
	// mapMake1()
	// mapMake2()
	//mapMake3()
	
	//最佳实践.
	//mapBestDemo()

	//map的增删改查
	//mapCrud()

	//map的遍历
	// mapIterate()

	//map切片:即切片的数据类型是map. 主要用来动态增加map内容.
	//mapSlice()



	

	

}