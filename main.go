package main

import (
	"github.com/guosenor/web-spider/engine"
	"github.com/guosenor/web-spider/scheduler"
	"github.com/guosenor/web-spider/zhenai/parser"
)

func main() {
	//engine.SimpleEngine{}.Run(engine.Request{
	//	Url:"http://city.zhenai.com/",
	//	ParserFunc:parser.ParserCityList,
	//})

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.SampleScheduler{},
		WorkerCount: 50,
	}
	e.Run(engine.Request{
		Url:        "http://city.zhenai.com/",
		ParserFunc: parser.ParserCityList,
	})
}
