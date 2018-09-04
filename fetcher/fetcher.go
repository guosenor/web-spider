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
	"log"
	"time"
)

var limter  = time.Tick(100 * time.Millisecond)
func Fetch(url string)([]byte,error)  {

	<-limter
	log.Printf("Fecth url:%s  \n",url)

	httpClient:=&http.Client{}
	request,_ := http.NewRequest("GET",url,nil)
	request.Header.Add("User-Agent","Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36")
	res, err := httpClient.Do(request)

	if err != nil {
		return nil ,err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil ,fmt.Errorf(" wrong statusCode:%d",res.StatusCode)
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