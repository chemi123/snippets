package main

import (
	"fmt"
)

func main() {
	set := make(map[string]struct{})
	set["set"] = struct{}{}
	fmt.Println(set)
}
