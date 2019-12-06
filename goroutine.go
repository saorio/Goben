package main

import (
    "fmt"
    "time"
)

func threeSecond() {
    time.Sleep(3000 * time.Millisecond)
    fmt.Println("3 seconds passed")
}

func oneSecond() {
    time.Sleep(1000 * time.Millisecond)
    fmt.Println("1 second passed")
}

func twoSecond() {
    time.Sleep(2000 * time.Millisecond)
    fmt.Println("2 second passed")
}

func main() {

    fmt.Println(time.Now())
    go threeSecond()
    go oneSecond()
    go twoSecond()
    fmt.Println(time.Now())
}
