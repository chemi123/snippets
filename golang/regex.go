package main

import (
	"fmt"
	"regexp"
)

func main() {
	r := regexp.MustCompile("hoge")
	if r.MatchString("fugahogepiyo") {
		fmt.Println("matched")
	}
}
