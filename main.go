package main

import (
	"github.com/ChimeraCoder/anaconda"
	"fmt"
//	"time"
)

var (
	ConsumerKey string
	ConsumerSecret string
	AccessToken string
	AccessTokenSecret string
)

func main() {
	ConsumerKey := "iaQcKNAy0ctYb4gIIfji7xSNR"
	ConsumerSecret := "FEsuiDoCs6ftWAqMyLdFzSOGW1oixK5XMaYaoff7GNfT0dsGF3"
	AccessToken := "53841728-EgOzg31Ew7l0wUBlUtJFckW8usCskJxu98VPq5ACT"
	AccessTokenSecret := "mbDS9AfNApOHxQ2l20L3PAzxpqujTfdNA3zXlVGme7XTx"

	anaconda.SetConsumerKey(ConsumerKey)
	anaconda.SetConsumerSecret(ConsumerSecret)
	api := anaconda.NewTwitterApi(AccessToken, AccessTokenSecret)

	search_Result, _ := api.GetSearch("golang", nil)
	for _, tweet := range search_Result.Statuses {
		fmt.Println(tweet.Text)
		// api.SetDelay(10*time.Second)
	}
}