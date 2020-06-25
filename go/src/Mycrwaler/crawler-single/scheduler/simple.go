package scheduler

import (
	"Mycrwaler/crawler-single/engine"
	"log"
)


type SimpleScheduler struct {
	WorkerChan chan engine.Request
}

func (s *SimpleScheduler) Run() {
	s.WorkerChan = make(chan engine.Request)
}

func (s *SimpleScheduler) GetWorkerChan() chan engine.Request {
	return s.WorkerChan
}



func (s *SimpleScheduler) Submit(request engine.Request) {
	go func() {//TODO 如果 out队列不加缓冲1024的话  这里若不go   则会卡在fetching
		s.WorkerChan <- request
	}()
	log.Println(request.Url)

}



func (s *SimpleScheduler) WorkerReady(chan engine.Request) {

}

