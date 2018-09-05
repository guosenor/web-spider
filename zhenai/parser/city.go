package parser

import (
	"github.com/guosenor/web-spider/engine"
	"regexp"
)
var cityRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
var cityUrlRe = regexp.MustCompile(`(http://www.zhenai.com/zhenghun/[^"]+)`)

func ParseCtiy(contents[]byte)engine.ParserResult  {

	matches := cityRe.FindAllSubmatch(contents, -1)
	result:=engine.ParserResult{}
	for _, m:=range matches{
		//result.Items= append(result.Items,"User "+string(m[2]))
		//log.Printf("get User %s \n",string(m[2]))
		result.Requests=append(result.Requests,engine.Request{
			Url:string(m[1]),
			ParserFunc:ParserUserProfile,
		})
	}
	matches = cityUrlRe.FindAllSubmatch(contents,-1)
	for _, m:=range matches{
		//result.Items= append(result.Items,"City "+string(m[1]))
		result.Requests=append(result.Requests,engine.Request{
			Url:string(m[1]),
			ParserFunc:ParseCtiy,
		})
	}
	return result
}
