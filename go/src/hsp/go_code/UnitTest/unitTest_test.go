package main

import (
	_ "fmt"
	"testing"
)
// 测试案例怎么写?


// testing框架的使用:
// 1,测试文件必须以_test.go结尾,测试用例函数必须Test打头
// 2,命令行运行go test -v,  
// 3,也可以指定测试某个文件go test -v unitTest_test.go unitTest.go
// 4,指定测试某个方法go test -v -test.run TestAddMethod
func TestAddMethod(t *testing.T){

	res := addMethod(20,30)
	if res != 50{
		// fmp.Println("测试案例错误")
		t.Fatalf("执行错误,期望值应该是%v,实际值是%v",50,res)
	}
	t.Logf("执行正确addMethod")
}

func TestHello(t *testing.T){
	t.Logf("执行正确hello")
}