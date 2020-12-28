package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

func main() {
	//requestGetBilety()
	requestGET()
	requestPOST()
	customRequest()
	postingAFormAndDecodingJSON()
}

func postingAFormAndDecodingJSON() {
	formData := url.Values{
		"name": {"marek nieznany"},
	}

	resp, err := http.PostForm("https://httpbin.org/post", formData)
	if err != nil {
		log.Fatalln(err)
	}

	var res map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&res)

	log.Println(res["form"])
}

func customRequest() {
	requestBody, err := json.Marshal(map[string]string{
		"name":  "marek nieznany",
		"email": "marek@o2.pl",
	})
	if err != nil {
		log.Fatalln(err)
	}

	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	request, err := http.NewRequest("POST", "https://httpbin.org/post", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-type", "application/json")

	if err != nil {
		log.Fatalln(err)
	}
	resp, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("\n Custom request:")
	log.Println(string(body))
}

func requestPOST() {
	requestBody, err := json.Marshal(map[string]string{
		"name":  "leon nieznany",
		"email": "leon@o2.pl",
	})
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(body))
}

func requestGET() {
	resp, err := http.Get("https://httpbin.org/cookies")
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(body))
}

func requestGetBilety() {
	resp, err := http.Get("http://bilety-czarterowe-test.r.pl:8516/pobierzPrzeloty")
	if err != nil {
		//return nil, err
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(body))
}
