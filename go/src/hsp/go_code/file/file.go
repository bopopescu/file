package main

import (
	"fmt"
	"os"
	"io"
	"bufio"
)

func readFile(){
	fielPath := "D:/read.txt"
	// 判断文件目录是否存在
	_.error :=  os.Stat(fielPath)
	if error != nil{
		return
	}
	file,err := os.OpenFile(fielPath,os.O_RDONLY,2524)
	if err != nil{
		fmt.Println("打开文件错误",err)
	}
	// 当函数退出时,关闭
	defer file.Close()
	// 默认4096的缓存
	reader := bufio.NewReader(file)
	// 循环读取
	for{
		str,err := reader.ReadString('\n') // 读到一个换行就结束
		if err == io.EOF{   //io.EOF代码文件末尾
			break
		}
		fmt.Print(str)
	}
}
func writeFileDemo1(){
	// 创建一个新文件
	fielPath := "D:/write.txt"

	// 没有文件就创建, O_TRUNC 先清空 ,再写入新的内容
	// O_APPEND  追加
	file,err := os.OpenFile(fielPath,os.O_WRONLY|os.O_CREATE|os.O_APPEND,2524)
	if err != nil{
		fmt.Println("打开文件错误",err)
		return;
	}
	defer file.Close()
	// 写入
	str := "hello,dawsen \r\n"
	writer := bufio.NewWriter(file)
	for i:=0;i<5;i++{
		writer.WriteString(str)
	}

	writer.Flush()

}
func readAndWrite(){
	// 创建一个新文件
	fielPath := "D:/write.txt"
	// 没有文件就创建, O_TRUNC 先清空 ,再写入新的内容
	// O_APPEND  追加
	file,err := os.OpenFile(fielPath,os.O_RDWR|os.O_CREATE|os.O_APPEND,0666)
	if err != nil{
		fmt.Println("打开文件错误",err)
		return;
	}
	defer file.Close()

	// 读取原来的内容
	read := bufio.NewReader(file)
	for{
		str,err := read.ReadString('\n')
		if err == io.EOF{
			break
		}
		fmt.Print(str)
	}


	// 写入
	str := "hello,dawsen \r\n"
	writer := bufio.NewWriter(file)
	for i:=0;i<5;i++{
		writer.WriteString(str)
	}

	writer.Flush()
}
func main(){
	// 读取文件内容. file是一个文件指针(句柄)
	// readFile()

	// 写文件
	// writeFileDemo1()


	// 边读边写.
	readAndWrite()


	// os.Args可以拿到命令行输入的所有参数, flag包也可以获取到命令行参数.
}