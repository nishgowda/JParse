package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	jparse "github.com/nishgowda/JParse"
)

func generateURL() string {
<<<<<<< HEAD:Tests/Test-2/test-2.go
	key := "KEY" // -> Edit key out in production
=======
	key := "KEY"
>>>>>>> 8597a1491502bf8ef6a34363abd890c0724fb8fa:testing/test_2/test_2.go
	req, _ := http.NewRequest(
		"GET",
		"http://dataservice.accuweather.com/forecasts/v1/hourly/12hour/3193_PC",
		nil)
	q := req.URL.Query()
	q.Add("apikey", key)
	q.Add("metric", "true")
	q.Add("details", "true")
	req.URL.RawQuery = q.Encode()
	return req.URL.String()
}

func main() {
	url := generateURL()
	//fmt.Println(url)
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	js := string(body)
	values := []string{"DateTime", "IconPhrase"}
	embObj := []string{"Temperature"}
	embVal := []string{"Value", "Unit"}
	parsed := jparse.EmbeddedObjArrayParse(values, js, embObj, embVal)
	fmt.Println(parsed)

}
