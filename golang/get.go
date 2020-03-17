package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "http://sample.jp/"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(bytes))
}
