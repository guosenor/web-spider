package main

import (
	"github.com/guosenor/web-spider/engine"
	"github.com/guosenor/web-spider/scheduler"
	"github.com/guosenor/web-spider/zhenai/parser"
	"github.com/guosenor/web-spider/persist"
)

func main() {

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:persist.ItemSave(),
	}
	e.Run(engine.Request{
		//Url:        "http://city.zhenai.com/",
		//ParserFunc: parser.ParserCityList,
		Url:"http://www.zhenai.com/zhenghun",
		ParserFunc:parser.ParseCtiy,
	})
}