package parser

import (
	"testing"
	"io/ioutil"
)

func TestParserCityList(t *testing.T) {
	contents, err := ioutil.ReadFile("cityList.html")
	if err != nil {
		panic(err)
	}
	result := ParserCityList(contents)
	const size = 470
	expectedUrls := []string{"http://city.zhenai.com/aba", "http://city.zhenai.com/akesu", "http://city.zhenai.com/alashanmeng"}
	expectedCities := []string{"阿坝", "阿克苏", "阿拉善盟"}

	if len(result.Requests) != size {
		t.Error("result should have %d Requests but had %d", size, len(result.Requests))
	}
	for i, url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Error("expected Url #%d: %s; but is %s", i, result.Requests[i].Url)
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
