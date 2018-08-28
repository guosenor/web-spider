package engine

import (
	"github.com/guosenor/web-spider/fetcher"
	"log"
)

func Run(seeds ...Request)  {
	var requests []Request
	for _,r :=range seeds  {
		requests=append(requests,r)
	}
	for len(requests)>0{
		r:=requests[0]
		log.Printf("Url:%s\n",r.Url)
		requests = requests[1:]
		body,err:=fetcher.Fetch(r.Url)
		if err!=nil{
			log.Print("Fetcher error fetch url: %s %v \n",r.Url,err)
			continue
		}
		parseResult := r.ParserFunc(body)
		requests = append(requests,parseResult.Requests...)
	}
}
