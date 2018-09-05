package persist

import (
	"log"
	"gopkg.in/olivere/elastic.v5"
	"context"
)

func ItemSave() chan interface{} {
	out := make(chan interface{})
	itemCount := 0
	go func() {
		for {
			item := <-out
			itemCount++
			log.Printf("item save got a item %d -- %v \n", itemCount, item)
		}
	}()
	return out
}
func saveToElastic(item interface{}) (id string, err error) {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return "", err
	}
	res, err := client.Index().Index("dating_profile").Type("zhenai").BodyJson(item).Do(context.Background())
	if err != nil {
		return "",err
	}
	return res.Id,nil
}
