package main

import (
	"fmt"
	"github.com/karishmashuklaa/hc/httpcl"
)

func main() {
	client := httpcl.New()

	response, err := client.Get("https://api.github.com", nil)

	if err != nil {
		panic(err)
	}

	fmt.Println(response.StatusCode)
}