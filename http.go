package main

import (
	"fmt"
	"net/http"
)

func requets_test() {
	url := "https://google.com"
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	fmt.Println(resp.StatusCode)

	// byteArray, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(byteArray))
}
