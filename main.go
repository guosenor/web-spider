package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
)

func main() {
	res, err := http.Get("http://city.zhenai.com/")

	if err != nil {
		panic("err")
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return
	}
	e := determineEncoding(res.Body)
	utf8Reader := transform.NewReader(res.Body, e.NewDecoder())
	all, err := ioutil.ReadAll(utf8Reader)
	printCitiList(all)
}
func determineEncoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
func printCitiList(contents []byte) {
	re := regexp.MustCompile(`<a href="(http://city.zhenai.com/[0-9a-z]+)"[^>]*>([^<]+)</a>`)
	matches := re.FindAllSubmatch(contents, -1)
	for _, m := range matches {

		fmt.Printf("cityName: %s    Url: %s\n", m[2], m[1])
	}
	fmt.Printf("city count:%d\n", len(matches))
}
