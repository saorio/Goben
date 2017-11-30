package main

import (
	"fmt"

	"github.com/bradfitz/gomemcache/memcache"
)

func main() {
	mc := memcache.New("127.0.0.1:11211")
	mc.Set(&memcache.Item{Key: "Name", Value: []byte("Saori")})

	it, err := mc.Get("Name")
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(it.Key, it.Value)
}
