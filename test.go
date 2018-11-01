package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "http://google.com"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	text, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)
	fmt.Println(string(text))
}
