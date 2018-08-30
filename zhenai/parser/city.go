package parser

import (
	"github.com/guosenor/web-spider/engine"
	"regexp"
	"log"
)
const cityRe=`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

func ParseCtiy(contents[]byte)engine.ParserResult  {
	re := regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(contents, -1)
	result:=engine.ParserResult{}
	for _, m:=range matches{
		result.Items= append(result.Items,"User "+string(m[2]))
		log.Printf("get User %s \n",string(m[2]))
		result.Requests=append(result.Requests,engine.Request{
			Url:string(m[1]),
			ParserFunc:ParserUserProfile,
		})
	}
	return result
}
