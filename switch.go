package main

import (
	"fmt"
)

func main() {
	var i int
	for i = 1; i <= 100; i++ {
		var a int
		a = i % 2
		switch a {
		case 0:
			fmt.Println(i, "数値-偶数")
		case 1:
			fmt.Println(i, "数値-奇数")
		}
	}
}

