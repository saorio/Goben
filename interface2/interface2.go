package main

import "fmt"

// Prefecture is struct
type Prefecture struct {
	Number string
	Name   string
}

// Call return string
func (p *Prefecture) Call() string {
	return p.Number + p.Name
}

// CallPrefecture describeというものを以下で定義、string型で返すCallメソッドを1つだけ持つ
type CallPrefecture interface {
	Call() string
}

func printPrefecture(print CallPrefecture) {
	fmt.Println(print.Call())
}

func main() {
	prefecture := &Prefecture{"110-0014", "台東区"}
	printPrefecture(prefecture)
}
