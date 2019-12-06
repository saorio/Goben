package main

import (
	"fmt"

	"github.com/ChimeraCoder/anaconda"
	//	"time"
)

// ConsumerKey key
var ConsumerKey string

// ConsumerSecret secret
var ConsumerSecret string

// AccessToken token
var AccessToken string

// AccessTokenSecret secret
var AccessTokenSecret string

func main() {
	ConsumerKey := ""
	ConsumerSecret := ""
	AccessToken := ""
	AccessTokenSecret := ""

	anaconda.SetConsumerKey(ConsumerKey)
	anaconda.SetConsumerSecret(ConsumerSecret)
	api := anaconda.NewTwitterApi(AccessToken, AccessTokenSecret)

	searchResult, _ := api.GetSearch("golang", nil)
	for _, tweet := range searchResult.Statuses {
		fmt.Println(tweet.Text)
		// api.SetDelay(10*time.Second)
	}
}
