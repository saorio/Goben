package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Print(rand.Intn(10), " ")
	}
	fmt.Println(" ")

	myrand := rand.New(rand.NewSource(1))
	for i := 0; i < 10; i++ {
		fmt.Print(myrand.Intn(10), " ")
	}
	fmt.Println(" ")

	myrand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		fmt.Print(myrand.Intn(10), " ")
	}
}
