package engine

// ConcurrentEngine 并发引擎
type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
	ItemChan    chan interface{}
}

// Scheduler 调度
type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkerChan() chan Request
	Run()
}

// ReadyNotifier 准备
type ReadyNotifier interface {
	WorkerReady(chan Request)
}

// Run 运行
func (e *ConcurrentEngine) Run(seeds ...Request) {
	gotUrls := make(map[string]int)
	out := make(chan ParserResult)
	e.Scheduler.Run()
	for i := 0; i < e.WorkerCount; i++ {
		createWorker(e.Scheduler, e.Scheduler.WorkerChan(), out)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	for {
		result := <-out
		for _, item := range result.Items {
			go func() {
				e.ItemChan <- item
			}()
		}
		//fmt.Println()

		for _, request := range result.Requests {
			if gotUrls[request.URL] == 0 {
				e.Scheduler.Submit(request)
				gotUrls[request.URL] = 1
			}

		}
	}
}
func createWorker(ready ReadyNotifier, in chan Request, out chan ParserResult) {
	go func() {
		for {
			ready.WorkerReady(in)
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
