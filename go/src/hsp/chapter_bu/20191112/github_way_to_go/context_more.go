package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	//ctx, cancel := context.WithTimeout(context.TODO(), time.Second*3)
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*1)

	defer cancel() // 防止任务比超时时间短导致资源未释放

	// 启动协程
	go task(ctx)

	// 主协程需要等待，否则直接退出
	time.Sleep(time.Second * 4)

}

func task(ctx context.Context) {

	ch := make(chan struct{})

	// 真正的任务协程
	go func() {
		// 模拟两秒耗时任务
		time.Sleep(time.Second * 2)
		ch <- struct{}{}
	}()

	for {
		select {
		case k, ko := <-ch:
			fmt.Println("done", k, ko)
			return
		case i, ok := <-ctx.Done():
			if ok {
				fmt.Println("timeout", i, ok)
				return
			} else {
				//说明 context的 给关闭掉了
				fmt.Println("timeout22", i, ok)

			}
		default:
			//fmt.Println("lvelvelve")
		}
	}

}
