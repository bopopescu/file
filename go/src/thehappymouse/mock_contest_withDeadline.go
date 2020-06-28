package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"
)

var (
	logg *log.Logger
)

func work(ctx context.Context, ch chan bool) {
	for {
		select {
		case <-ctx.Done():
			logg.Println(`下班!`)
			time.Sleep(8e9)

			ch <- true
			logg.Println(`111111111111111`)

			return
		default:
			logg.Println(`上班!`)
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	ch := make(chan bool)
	logg = log.New(os.Stdout, "", log.Ltime)
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
	go work(ctx, ch)


	time.Sleep(10 * time.Second)
	//取消函数：当cancel被调用时,context.WithDeadline设置的时间超过了，关闭ctx.Done通道。
	cancel()
	// 这个chan是为了保证子的goroutine执行完,当然也可以不用chan用time.Sleep停止几秒
	vivo :=<-ch
	fmt.Println(vivo)

	logg.Println(`无脑发呆中!`)
}
/* outfile:
17:29:43 上班!
17:29:45 上班!
17:29:47 上班!
17:29:49 下班!
17:29:53 无脑发呆中!
*/