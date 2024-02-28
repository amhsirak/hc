package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	url := "https://api.github.com"

	client := http.Client{}

	request, err := http.NewRequest("GET", url, nil)
	request.Header.Set("Accept", "application/xml")

	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.StatusCode)
	fmt.Println(string(bytes))
}