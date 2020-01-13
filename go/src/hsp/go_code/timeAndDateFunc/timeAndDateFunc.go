package main

import (
	"fmt"
	"time"
)

func main() {
	// 1,获取当前时间
	now := time.Now()
	fmt.Printf("now的值是%v,now的类型是%T", now, now)

	// 2,上面获取的now的类型是一个结构体
	fmt.Printf("now的年值是%v", now.Year())

	// 3,格式化日期第一种方式
	dateStr := fmt.Sprintf("当前时间是: %d-%d-%d %d:%d:%d", now.Year(),
		now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Println("", dateStr)

	// 4,格式化日期第二种方式
	fmt.Printf(now.Format("2006/01/02 15:04:05"))

	// 5,最佳实践:每隔一秒钟,打印一个数字,直到100的时候退出.
	// 使用time包里面的常量
	// for i:=0;;i++{
	// 	fmt.Println("",i)
	// 	time.Sleep(time.Second)
	// 	if i>10 {
	// 		break
	// 	}
	// }

	// 6, unix时间戳
	fmt.Printf("now的秒时间戳是%v \t now的纳秒时间戳是%v", now.Unix(), now.UnixNano())
}
