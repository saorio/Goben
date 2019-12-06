package main

import (
	"fmt"
	"strconv"
)

// MyCar is structure
type MyCar struct {
	name  string
	speed int
}

// method
func (u *MyCar) run(speed int) string {
	u.speed = speed
	return strconv.Itoa(speed) + "kmで走ります"
}

func (u *MyCar) stop() {
	fmt.Println("停止します")
	u.speed = 0
}

func main() {
	myCar := &MyCar{name: "マイカー", speed: 0}

	var i Car = myCar
	fmt.Println(i.run(50))
	i.stop()
}
