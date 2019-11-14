package main

import (
	"web-spider/engine"
	"web-spider/persist"
	"web-spider/scheduler"
	"web-spider/zhenai/parser"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:    persist.ItemSave(),
	}
	e.Run(engine.Request{
		//URL:        "http://city.zhenai.com/",
		//ParserFunc: parser.ParserCityList,
		URL:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCtiy,
	})
}
