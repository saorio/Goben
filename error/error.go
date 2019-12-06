package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	if err := MyError(); err != nil {
		fmt.Println("error", err)
		os.Exit(1)
	}
	fmt.Println("Nothing happens.")
}
func MyError() error {
	return errors.New("My error occured.")
}
