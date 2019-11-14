package parser

import (
	"regexp"
	"web-spider/engine"
)

const cityListRe = `<a href="(http://city.zhenai.com/[0-9a-z]+)"[^>]*>([^<]+)</a>`

// ParserCityList 城市解析
func ParserCityList(contents []byte) engine.ParserResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParserResult{}
	for _, m := range matches {
		//result.Items= append(result.Items,"City "+string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			URL:        string(m[1]),
			ParserFunc: ParseCtiy,
		})
	}
	return result
}
