package main

// by default, content will be json.
// if you want to change the content type, you can do so by setting the content type header
// we do not want to force our users to marshal json themselves, so we will do it for them
// custom content type management.

import (
	"fmt"
	"github.com/karishmashuklaa/hc/httpcl"
	"io"
	"net/http"
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

type User struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
}

func getUrls() {
	response, err := githubHttpClient.Get("https://api.github.com", nil)

	if err != nil {
		panic(err)
	}

	fmt.Println(response.StatusCode)

	bytes, _ := io.ReadAll(response.Body)
	fmt.Println(string(bytes))
}

func createUser(user User) {
	response, err := githubHttpClient.Post("https://api.github.com", nil, user)
	
	if err != nil {
		panic(err)
	}

	fmt.Println(response.StatusCode)

}