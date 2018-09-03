package engine

import (
	"fmt"
	"log"
)

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}
type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkerChan() chan Request
	Run()
}
type ReadyNotifier interface {
	WorkerReady(chan Request)
} 
func (e *ConcurrentEngine) Run(seeds ...Request) {

	out := make(chan ParserResult)
	e.Scheduler.Run()
	for i := 0; i < e.WorkerCount; i++ {
		createWorker(e.Scheduler,e.Scheduler.WorkerChan(), out)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	for {
		result := <-out
		for _, item := range result.Items {
			log.Printf("Get item : %v", item)
		}
		fmt.Println()

		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}
func createWorker(ready ReadyNotifier,in chan Request,out chan ParserResult) {
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