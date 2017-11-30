package main

import (
	"fmt"

	"github.com/ChimeraCoder/anaconda"
	//	"time"
)

var (
	ConsumerKey       string
	ConsumerSecret    string
	AccessToken       string
	AccessTokenSecret string
)

func main() {
	ConsumerKey := ""
	ConsumerSecret := ""
	AccessToken := ""
	AccessTokenSecret := ""

	anaconda.SetConsumerKey(ConsumerKey)
	anaconda.SetConsumerSecret(ConsumerSecret)
	api := anaconda.NewTwitterApi(AccessToken, AccessTokenSecret)

	search_Result, _ := api.GetSearch("golang", nil)
	for _, tweet := range search_Result.Statuses {
		fmt.Println(tweet.Text)
		// api.SetDelay(10*time.Second)
	}
}
