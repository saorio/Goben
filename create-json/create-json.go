package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type Department struct {
	Name   string    `json:"name"`
	Member []Profile `json:"profile"`
}

type Profile struct {
	Name string `json:"name"`
	Age  string `json:"age"`
	Sex  string `json:sex"`
}

func main() {
	saori := Profile{Name: "saori", Age: "27", Sex: "female"}
	norikazu := Profile{Name: "norikazu", Age: "31", Sex: "male"}
	yosuke := Profile{Name: "yosuke", Age: "32", Sex: "male"}
	development := Department{
		Name:   "開発部",
		Member: []Profile{saori, norikazu, yosuke},
	}
	jsonBytes, err := json.Marshal(development)
	if err != nil {
		fmt.Println("JSON Mershal error:", err)
		return
	}
	formated := new(bytes.Buffer)
	json.Indent(formated, jsonBytes, "", "	")
	fmt.Println(formated.String())
}
