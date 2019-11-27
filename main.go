package main

import (
	"flag"
	"fmt"
)

const (
	RW_MEMORY_RSS int = 0
	RW_MEMORY_SHR int = 1
	RW_MEMORY_SLB int = 2
)

type progma_info struct {
	Author         string `json:"ryuichi1208"`
	Email          string `json:"ryucrosskey@gmail.com"`
	ProgramName    string `json:"mem-checker"`
	ProgramVersion string `json:"v0.0.1"`
}

type option_flag struct {
	mem_type int
}

// func requets_test() {
// 	url := "https://google.com"
// 	resp, _ := http.Get(url)
// 	defer resp.Body.Close()
// 	fmt.Println(resp.StatusCode)

// 	// byteArray, _ := ioutil.ReadAll(resp.Body)
// 	// fmt.Println(string(byteArray))
// }

func main() {
	fmt.Println("vim-go")

	var (
		i   = flag.Int("int", 0, "int flag")
		s   = flag.String("str", "default", "string flag")
		b   = flag.Bool("bool", false, "bool flag")
		url = flag.String("url", "google.com", "request url")
	)

	flag.Parse()
	fmt.Println(*i, *s, *b, *url, flag.NArg(), flag.NFlag())

	if *b {
		requets_test()
	}
}
