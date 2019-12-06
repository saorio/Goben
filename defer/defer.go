package main

import (
	"fmt"
)

func greeting() {
	defer func() {
		fmt.Println("Good Bye")
		err := recover()
		if err != nil {
			fmt.Println("recover system:", err)
		}
	}()

	fmt.Println("Hello")
	panic("panic system.")
}

func main() {
	greeting()
}
