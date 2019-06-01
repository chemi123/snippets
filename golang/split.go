package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "a:b:c:d"
	fmt.Println(strings.Split(str, ":"))
}
