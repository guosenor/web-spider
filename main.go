package main

import (
	"github.com/guosenor/web-spider/engine"
	"github.com/guosenor/web-spider/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:"http://city.zhenai.com/",
		ParserFunc:parser.ParserCityList,
	})
}

