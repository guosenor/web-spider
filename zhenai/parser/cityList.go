package parser

import (
	"github.com/guosenor/web-spider/engine"
	"regexp"
)
const cityListRe=`<a href="(http://city.zhenai.com/[0-9a-z]+)"[^>]*>([^<]+)</a>`
func ParserCityList(contents [] byte)engine.ParserResult  {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)
	result:=engine.ParserResult{}
	for _, m:=range matches{
		result.Items= append(result.Items,string(m[2]))
		result.Requests=append(result.Requests,engine.Request{
			Url:string(m[1]),
			ParserFunc:engine.NilPaser,
		})
	}
    return result
}