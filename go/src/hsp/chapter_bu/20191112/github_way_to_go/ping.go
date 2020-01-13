package main

import (
	"fmt"
	"net/http"
	"time"
)

var urls = []string{
	"http://www.baidu.com/",
	"http://www.clearloveuzi.cn/",
	"http://www.sina.com/",
	"http://www.qq.com/",
	"http://www.cctv.com/",
	"http://www.cnblog.cn/",
}

func main() {
	// Execute an HTTP HEAD request for all url's
	// and returns the HTTP status string or an error string.
	stime := time.Now().UnixNano()

	chan1 := make(chan bool, 10)
	//chan1:=make(chan bool)
	//_,ok:=<-chan1
	//
	//fmt.Println(ok)
	//return

	for _, url := range urls {
		go func(url string) {
			stime1 := time.Now().UnixNano()

			fmt.Println(stime1, "begin", url)
			resp, err := http.Head(url)
			if err != nil {
				fmt.Println("Error:", url, err)
			}
			etime1 := time.Now().UnixNano()

			fmt.Println(url, ": ", resp.Status, stime1, "--", (etime1-stime1)/1e6)

			fmt.Println("ssssss")
			chan1 <- true
			//close(chan1)
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		_, ok := <-chan1

		fmt.Println(ok, 6666666)
	}

	//<-chan1
	//<-chan1
	//<-chan1
	//<-chan1//todo 多一个就不会退出程序了哦  会一直阻塞   所以 还是相当于老韩笔记里的 for循环pop队列
	//todo important 注意 ！ 这里也不能用go协程单独运行哦  必须在主程序阻塞这个队列

	//			for i:=range chan1{  //todo  range 会一直阻塞住
	//	fmt.Println(i)
	//}

	etime := time.Now().UnixNano()
	fmt.Println("allfuck", (etime-stime)/1e6)

}
