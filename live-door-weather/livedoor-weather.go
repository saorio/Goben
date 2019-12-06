package main

import (
	"fmt"
)

// Prefecture is struct
type Prefecture struct {
	Number string
	Name   string
}

func (p *Prefecture) CallPrefecture() string {
	return p.Number + p.Name
}

type District interface {
	CallPrefecture() string
}

func describePrefecture(district District) {
	fmt.Println(district.District)
}
func main() {
	Prefecture := &Prefecture{"110-0014", "台東区北上野"}
	ptintName(Prefecture)
}
