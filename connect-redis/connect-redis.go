package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/garyburd/redigo/redis"
)

func main() {
	var h = flag.String("hostname", "127.0.0.1", "Set Hostname")
	var p = flag.String("port", "6379", "Set Port")
	var key = flag.String("key", "Name", "Set Key")
	var val = flag.String("value", "saorio", "Set Value")
	flag.Parse()

	c, err := redis.Dial("tcp", *h+":"+*p)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer c.Close()

	c.Do("SET", *key, *val)
	s, err := redis.String(c.Do("GET", *key))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("%#v\n", s) // ダブルクォートの改行で返す
}
