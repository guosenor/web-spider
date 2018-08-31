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
func (s *SampleScheduler) ConfigureMasterWorkerChan( c chan engine.Request)  {
	s.workerChan=c
}