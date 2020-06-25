package engine

import (
	"fmt"

)
//TODO 和done.go  同步看问题
// 并发的引擎
// 引擎将请求发送给调度器，调度器纷发给workers, workers的结果再返回给引擎
// 所有worker 引用一个源
type ConcurrentEngine struct {
	MaxWorkerCount int
	Scheduler      Scheduler
	RequestWorker Processor
	ItemChan       chan Item

}
type Processor func(request Request) (ParseResult, error)


type Scheduler interface {
	Submit(request Request)
	GetWorkerChan() chan Request


	Run()
	Ready
}

type Ready interface {
	WorkerReady(chan Request)
}



func (e *ConcurrentEngine) Run(seed ...Request) {
	//out := make(chan ParseResult, 1024)
	out := make(chan ParseResult)

	e.Scheduler.Run()

	for i := 0; i < e.MaxWorkerCount; i++ {
		e.createWorker(e.Scheduler.GetWorkerChan(), out, e.Scheduler)
	}

	for _, r := range seed {
		//if IsDuplicate(r.Url) {
		//	continue
		//}
		fmt.Println(666)
		e.Scheduler.Submit(r)
	}
	itemCount,testv := 0,0
	for {
		result := <-out

		for _, item := range result.Items {
			itemCount++
			fmt.Printf("Got Item: #%d %v", itemCount, item)
			//go func() { e.ItemChan <- item }()
		}

		testv++
		for _, r := range result.Requests {
			
			//if IsDuplicate(r.Url) {
			//	continue
			//}
			//fmt.Println(testv,"---",g,len(result.Requests)) //
				e.Scheduler.Submit(r)//todo 这里是为什么数量大于的时候不卡的关键   第10个的时候这里卡住了
				// todo  因为大于的时候  这个for循环会退出   进入上面的接受out大循环
			//fmt.Println("fuckin",g)//todo 正好第十个 注意 卡住了  为什么协程数量大于的时候不会卡死呢
		}
		//fmt.Println("aaaaaaaaaaaaaaaaaa")
	}

}
func (e *ConcurrentEngine) createWorker(in chan Request, out chan ParseResult, s Ready) {


	go func() {
		//s.WorkerReady(in) //TODO 为什么放这里就不行呢   queue的时候只能 读9个？？？
		//TODO 解  应该是因为 workchan 要重新放进去！！！   因为 workerQ = workerQ[1:]

		for {
			s.WorkerReady(in)
			request := <-in
			result, err := e.RequestWorker(request)
			if err != nil {
				continue
			}
			//fmt.Println("chuqu begin")
			out <- result
			//fmt.Println("chuqu end") //TODO  都没有打印这个 说明问题在out这   也就解释了为什么给out加了缓冲区就解决了问题
//TODO  同时为什么给in加了go func 也能解决呢？？？
		}
	}()
}

//举一反三
//TODO 1 协程数量大于 一页数据
//TODO 2   使用go func submit
////TODO 3    增加缓冲