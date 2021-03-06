package parser

import (
	"io/ioutil"
	"testing"
)

func TestParserCityList(t *testing.T) {
	contents, err := ioutil.ReadFile("cityList.html")
	if err != nil {
		panic(err)
	}
	result := ParserCityList(contents)
	const size = 470
	expectedUrls := []string{"http://city.zhenai.com/aba", "http://city.zhenai.com/akesu", "http://city.zhenai.com/alashanmeng"}
	expectedCities := []string{"City 阿坝", "City 阿克苏", "City 阿拉善盟"}

	if len(result.Requests) != size {
		t.Error("result should have %d Requests but had %d", size, len(result.Requests))
	}
	for i, url := range expectedUrls {
		if result.Requests[i].URL != url {
			t.Error("expected Url #%d: %s; but is %s", i, result.Requests[i].URL)
		}
	}
	if len(result.Items) != size {
		t.Error("result should have %d Items but had %d", size, len(result.Items))
	}
	for i, city := range expectedCities {
		if result.Items[i] != city {
			t.Error("expected Url #%d: %s; but is %s", i, result.Items[i])
		}
	}
}
