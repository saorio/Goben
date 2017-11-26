package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func main() {
	name := "bot"
	text := "Hello World"
	channel := "test"

	jsonStr := `{"channel":"` + channel + `","username":"` + name + `","text":"` + text + `"}`

	req, err := http.NewRequest(
		"POST",
		"https://hooks.slack.com/services/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",
		bytes.NewBuffer([]byte(jsonStr)),
	)

	if err != nil {
		fmt.Print(err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Print(err)
	}

	fmt.Print(err)
	defer resp.Body.Close()
}
