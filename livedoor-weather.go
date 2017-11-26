package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var RSS_FEED_TOKYO string = "http://weather.livedoor.com/forecast/rss/area/130010.xml"

type WeatherInfo struct {
	Title       string   `xml:"channel>title"`
	Description []string `xml:"channel>item>description"`
}

func main() {
	wi, err := DescribeWether(RSS_FEED_TOKYO)

	if err != nil {
		log.Fatal("Log: %v", err)
		return
	}
	fmt.Println(wi.Title)
	for n, v := range wi.Description {
		if n > 0 {
			fmt.Printf("%s \n", v)
		}
	}
}

func DescribeWether(feed string) (p *WeatherInfo, err error) {
	res, err := http.Get(feed)
	if err != nil {
		return nil, err
	}
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	wi := new(WeatherInfo)
	err = xml.Unmarshal(b, &wi)

	return wi, err
}
