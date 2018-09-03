package scheduler

import "github.com/guosenor/web-spider/engine"

type SampleScheduler struct {
	workerChan chan  engine.Request
}

func (s *SampleScheduler) Submit(r engine.Request)  {
	go func() {
		s.workerChan <- r
	}()
}
func (s *SampleScheduler) WorkerChan() chan engine.Request {
    return   s.workerChan
}
func (s *SampleScheduler) Run()  {
	s.workerChan = make(chan engine.Request)
}
func (s *SampleScheduler) WorkerReady(w chan engine.Request)  {
}