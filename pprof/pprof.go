package main

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"
)

func main() {
	cpuprofile := "mycpu.prof"
	f, err := os.Create(cpuprofile)
	if err != nil {
		log.Fatal("could not create CPU profile: ", err)
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()
	var i int
	for i = 1; i <= 100000; i++ {

		if i%2 == 0 {
			fmt.Println(i, "数値-偶数")
		} else {
			fmt.Println(i, "数値-奇数")
		}
	}
}
