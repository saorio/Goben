package main

import (
	"fmt"
)

func main() {
	var i int
	for i = 1; i <= 100; i++ {

		if i%2 == 0 {
			fmt.Println(i, "this is even.")
		} else {
			fmt.Println(i, "this is odd.")
		}
	}
}
