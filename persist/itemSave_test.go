package persist

import (
	"testing"
	"web-spider/model"
	"gopkg.in/olivere/elastic.v5"
	"context"
	"encoding/json"
)

func TestSaveToElastic(t *testing.T) {
	item := model.UserProfile{
		Name:"guosen",
		Gender:"男",
		Age:18,
	}
	id,err := saveToElastic(item)
	if err!= nil {
		panic(err)
	}
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	res,err := client.Get().Index("dating_profile").Type("zhenai").Id(id).Do(context.Background())
	if err!=nil {
		panic(err)
	}
	var actual model.UserProfile
	err = json.Unmarshal(*res.Source, &actual)
	if err !=nil {
		panic(err)
	}

	if item != actual {
		t.Errorf("got %v item %v" ,item,actual)
	}

}