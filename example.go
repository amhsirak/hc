package main

import (
	"fmt"
	"github.com/karishmashuklaa/hc/httpcl"
	"net/http"
)

func main() {
	getUrls()
}

func getUrls() {
	client := httpcl.New()

	headers := make(http.Header)
	headers.Set("Authorization", "Bearer ABC-123")

	response, err := client.Get("https://api.github.com", headers)

	if err != nil {
		panic(err)
	}

	fmt.Println(response.StatusCode)
}