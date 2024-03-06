package main

import (
	"fmt"
	"github.com/karishmashuklaa/hc/httpcl"
	"net/http"
	"io/ioutil"
)

var (
	githubHttpClient = getGithubClient()
)

func getGithubClient() httpcl.HttpClient {
	client := httpcl.New()

	commonHeaders := make(http.Header)
	commonHeaders.Set("Authorization", "Bearer ABC-123")
	client.SetHeaders(commonHeaders)

	return client
}

func main() {
	getUrls()
}

func getUrls() {
	response, err := githubHttpClient.Get("https://api.github.com", nil)

	if err != nil {
		panic(err)
	}

	fmt.Println(response.StatusCode)

	bytes, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(bytes))
}