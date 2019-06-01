package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("./tmp")

	if err != nil {
		fmt.Println(err)
		return
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	defer file.Close()

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
