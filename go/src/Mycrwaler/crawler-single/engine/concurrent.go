package engine

import "fmt"

type ConcurrentEngine struct {
	MaxWorkerCount int
	Scheduler      Scheduler
	ItemChan       chan Item
	RequestWorker Processor
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
		if IsDuplicate2(r.Url) {
			continue
		}
		fmt.Println(666)
		e.Scheduler.Submit(r)
	}


	itemCount,testv := 0,0

	for {
		result := <-out

		for _, item := range result.Items {
			itemCount++
			fmt.Printf("Got Item: #%d %v", itemCount, item)
			go func() { e.ItemChan <- item }()
		}

		testv++
		for _, r := range result.Requests {

			if IsDuplicate2(r.Url) {
				continue
			}
			e.Scheduler.Submit(r)//todo 这里是为什么数量大于的时候不卡的关键   第10个的时候这里卡住了
		}
	}

}

func (e *ConcurrentEngine) createWorker(in chan Request, out chan ParseResult, s Ready) {


	go func() {

		for {
			s.WorkerReady(in)
			request := <-in
			result, err := e.RequestWorker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
