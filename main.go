package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	url := ""

	client := http.Client{}
	response, err := client.Get(url)
	if err != nil {
		panic(err)
	}

	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.StatusCode)
	fmt.Println(string(bytes))
}