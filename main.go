package main

import (
	"flag"
	"fmt"
)

func main() {
	fmt.Println("hello world")

	src_url := flag.String("src_url", "", "")
	flag.Parse()
	fmt.Println(*src_url)
}
