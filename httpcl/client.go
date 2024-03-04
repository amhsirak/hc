package httpcl

import "net/http"

type httpClient struct{}

func New() HttpClient {
	client := &httpClient{}
	return client
}

type HttpClient interface{}

func (c *httpClient) Get(url string, headers http.Header) (*http.Response, error) {
	return c.do(method: "GET", url, headers, nil)
}

func Post(c *httpClient) {}

func Put(c *httpClient) {}

func Patch(c *httpClient) {}

func Delete(c *httpClient) {}