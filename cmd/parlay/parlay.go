package main

import (
	"flag"
	"fmt"
	"net/url"
	"os"

	"github.com/daetal-us/parlay"
)

func main() {
	flag.Parse()
	i := flag.Arg(0)
	if i == "" {
		fmt.Println("A path or URL is required.")
		os.Exit(1)
	}
	_, t := url.ParseRequestURI(i)
	var url string
	var err error
	if t != nil {
		url, err = parlay.FromPath(i)
	} else {
		url, err = parlay.FromURL(i)
	}
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(url)
}
