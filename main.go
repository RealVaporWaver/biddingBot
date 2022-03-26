package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	http "github.com/useflyent/fhttp"
)

func getPageInfo(slug string) *http.Response {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", fmt.Sprintf("https://opensea.io/collection/%s", slug), nil)
	req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36")
	res, _ := client.Do(req)

	return res
}

func removeDuplicateStr(strSlice []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

type Config struct {
	Address        string  `json:"address"`
	BidAmount      float32 `json:"bidAmount"`
	ExpirationTime int32   `json:"expirationTime"`
	AccountAddress string  `json:"accountAddress"`
	PrivateKey     string  `json:"privateKey"`
}

func main() {
	fmt.Println("Enter collection you want to bid on: ")
	var slug string
	fmt.Scanln(&slug)

	resp := getPageInfo(slug)

	doc, _ := goquery.NewDocumentFromReader(resp.Body)

	var a []string

	doc.Find("body a").Each(func(index int, item *goquery.Selection) {
		wed, _ := item.Attr("href")
		if strings.Contains(wed, "/assets/") {
			a = append(a, wed)
		}
	})

	a = removeDuplicateStr(a)
	contractAddress := strings.Split(a[1], "/")[2]
	//fmt.Printf("%v", a)
	println(contractAddress)

	fi, _ := os.ReadFile("config.json")

	jsonBlob := []byte(fi)

	config := Config{}

	err := json.Unmarshal(jsonBlob, &config)

	if err != nil {
		println("successfully loaded data")
	}

	config.Address = contractAddress

	configJson, _ := json.Marshal(config)

	ioutil.WriteFile("config.json", configJson, 0644)

	/*hash, no := http.Get("https://ja3er.com/json")

	body, test := ioutil.ReadAll(hash.Body)

	println(string(body), test, no)*/

}
