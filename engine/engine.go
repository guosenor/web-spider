package engine

import (
	"web-spider/fetcher"
	"log"
)

type SimpleEngine struct {

} 
func (e SimpleEngine)Run(seeds ...Request)  {
	var requests []Request
	for _,r :=range seeds  {
		requests=append(requests,r)
	}
	for len(requests)>0{
		r:=requests[0]
		requests = requests[1:]

		parseResult ,err := worker(r)
		if err!=nil{
			continue
		}

		requests = append(requests,parseResult.Requests...)
	}
}

func worker(r Request)(ParserResult ,error)  {

	body,err:=fetcher.Fetch(r.Url)
	if err!=nil{
		log.Printf("Fetcher error fetch url: %s %v \n",r.Url,err)
		return ParserResult{},err
	}
	return r.ParserFunc(body),nil
}
