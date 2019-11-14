package scheduler

import "web-spider/engine"

// SampleScheduler 基础调度
type SampleScheduler struct {
	workerChan chan engine.Request
}

// Submit 提交
func (s *SampleScheduler) Submit(r engine.Request) {
	go func() {
		s.workerChan <- r
	}()
}

// WorkerChan 清除
func (s *SampleScheduler) WorkerChan() chan engine.Request {
	return s.workerChan
}

// Run 运行
func (s *SampleScheduler) Run() {
	s.workerChan = make(chan engine.Request)
}

// WorkerReady 准备
func (s *SampleScheduler) WorkerReady(w chan engine.Request) {
}
