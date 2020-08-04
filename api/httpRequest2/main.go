package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	getHTTPRequest()
	postHTTPRequest()
}

func getHTTPRequest() {
	resp, err := http.Get("http://webcode.me")

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {

		panic(err)
	}

	fmt.Println(string(body))
}

func postHTTPRequest() {

	data := url.Values{
		"name":       {"John Doe"},
		"occupation": {"gardener"},
	}

	resp, err := http.PostForm("https://httpbin.org/post", data)

	if err != nil {
		panic(err)
	}

	var res map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&res)

	fmt.Println(res["form"])
}
