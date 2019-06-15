package main

import (
	"fmt"
)

func takeSlice(s []int) {
	fmt.Printf("\n===takeSlicei===\n")
	fmt.Printf("%p <- different address from main\n", &s)
	fmt.Println(&s[0])
	fmt.Println(&s[1])
	fmt.Println(&s[2])
}

func main() {
	s := []int{1, 2, 3}
	fmt.Println("===main===")
	fmt.Printf("%p\n", &s)
	fmt.Println(&s[0])
	fmt.Println(&s[1])
	fmt.Println(&s[2])
	takeSlice(s)
}
