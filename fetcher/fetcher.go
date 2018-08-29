package fetcher

import (
	"golang.org/x/text/encoding"
	"bufio"
	"golang.org/x/net/html/charset"
	"net/http"
	"golang.org/x/text/transform"
	"io/ioutil"
	"fmt"
	"golang.org/x/text/encoding/unicode"
)

func Fetch(url string)([]byte,error)  {
	res, err := http.Get(url)

	if err != nil {
		return nil ,err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil ,fmt.Errorf("wrong statusCode:%d",res.StatusCode)
	}
	bodyReader:=bufio.NewReader(res.Body)
	e := determineEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}
func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}