package main

import (
	"fmt"
	"encoding/json"
)
type Ghost struct{
	// tag使用 
	Name string `json:"name"`  //反射机制
	Age int `json:"age"`
	Sex string `json:"sex"`
}
func serialize(){
	ghost := Ghost{"zobie",100,"male"}
	// json序列化, 
	jsonstr,err := json.Marshal(ghost)
	if err != nil {
		fmt.Println("错误",err);
	}
	fmt.Println("序列化后的结果是",string(jsonstr));
}
func unserialize(){
	str := " {\"name\":\"zobie\",\"age\":100,\"sex\":\"male\"}"
	// json反序列化,
	var ghost  Ghost
	err := json.Unmarshal([]byte(str),&ghost)
	if err != nil {
		fmt.Println("错误",err);
	}
	fmt.Println("反序列化后结果是",ghost);
}
func main(){
	// serialize()
	unserialize()
}