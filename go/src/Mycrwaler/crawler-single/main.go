package main

import (
	"Mycrwaler/crawler-single/engine"
	"Mycrwaler/crawler-single/persist"
	"Mycrwaler/crawler-single/zhenai/parser"
	"Mycrwaler/crawler-single/scheduler"
)

// 单机，并发
func main() {


	var seed []engine.Request

	seed = []engine.Request{
		{
			Url:   "http://www.zhenai.com/zhenghun/wuhan",
			Parse: engine.NewFuncParser(parser.ParseCity, "ParseCity"),
		},

	}
	//e := engine.SimpleEngine{}	//单机


	//e:=engine.SimpleEngine{}

	//
	//e := engine.ConcurrentEngine{
	//	MaxWorkerCount:   15,
	//	Scheduler:        &scheduler.SimpleScheduler{},
	//	RequestWorker: engine.Worker,
	//}
	ee,_ :=persist.ItemSaver("myelas2")

	e := engine.ConcurrentEngine{
		MaxWorkerCount:   10,
		Scheduler:        &scheduler.QueuedScheduler{},
		RequestWorker: engine.Worker,
		ItemChan:	ee,
	}
	e.Run(seed...)
}
