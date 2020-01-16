package main

import (
	"github.com/thehappymouse/ccmouse/crawler/engine"
	"github.com/thehappymouse/ccmouse/crawler/persist"
	"github.com/thehappymouse/ccmouse/crawler/scheduler"
	"github.com/thehappymouse/ccmouse/crawler/zhengai/parser"
)

// 单机，并发
func main() {

	//itemChan, err := persist.ItemSaver("profiles")
	//if err != nil {
	//	panic(err)
	//}

	var seed []engine.Request

	seed = []engine.Request{
		{
			Url:   "http://www.zhenai.com/zhenghun/wuhan",
			Parse: engine.NewFuncParser(parser.ParseCity, "ParseCity"),
		},
		//{
		//	Url:"http://album.zhenai.com/u/1978653413",
		//	Parse:parser.NewProfileParser("sd"),
		//},
		//{
		//	Url:       "http://www.zhenai.com/zhenghun/henan",
		//	ParseFunc: parser.ParseCity,
		//},
		//{
		//	Url:   "http://www.zhenai.com/zhenghun",
		//	Parse: engine.NewFuncParser(parser.ParseCityList, "ParseCityList"),
		//},
		//{
		//	Url:   "http://cdn-imgbed.clearloveuzi.cn",
		//	Parse: en		gine.NewFuncParser(parser.Parseimgy, "Parseimgy"),
		//},
	}
	//e := engine.SimpleEngine{}	//单机

	//
	//e := engine.ConcurrentEngine{
	//	MaxWorkerCount:   10,
	//	Scheduler:        &scheduler.SimpleScheduler{},
	//	RequestWorker: engine.Worker,
	//}

	ee,_ :=persist.ItemSaver("myelas")
	e := engine.ConcurrentEngine{
		MaxWorkerCount:   10,
		Scheduler:        &scheduler.QueuedScheduler{},
		RequestWorker: engine.Worker,
		ItemChan: ee,
	}
	e.Run(seed...)
}
