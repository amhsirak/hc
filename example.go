package main

import (
	"fmt"
	"github.com/karishmashuklaa/hc/httpcl"
)

func main() {
	client := httpcl.New()

	headers := make(http.Header)
	headers.Set("Authorization", "Bearer ABC-123")

	response, err := client.Get("https://api.github.com", headers)

	if err != nil {
		panic(err)
	}

	fmt.Println(response.StatusCode)
}