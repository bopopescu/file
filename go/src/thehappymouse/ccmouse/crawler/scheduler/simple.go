package scheduler

import "github.com/thehappymouse/ccmouse/crawler/engine"

// 简单的并发处理
// 所有工作协程，共用一个输入通道
type SimpleScheduler struct {
	WorkerChan chan engine.Request
}


func (s *SimpleScheduler) GetWorkerChan() chan engine.Request {
	return s.WorkerChan
}

func (s *SimpleScheduler) Run() {
	s.WorkerChan = make(chan engine.Request)
}




func (s *SimpleScheduler) WorkerReady(chan engine.Request) {

}

func (s *SimpleScheduler) Submit(request engine.Request) {
	go func() {//TODO 如果 out队列不加缓冲1024的话  这里若不go   则会卡在fetching
		s.WorkerChan <- request
	}()
}
