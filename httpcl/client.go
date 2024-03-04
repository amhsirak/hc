package httpcl

import "net/http"

type httpClient struct{}

func New() HttpClient {
	client := &httpClient{}
	return client
}

type HttpClient interface{}

func (c *httpClient) Get(url string, headers http.Header) (*http.Response, error) {
	return c.do(http.MethodGet, url, headers, nil)
}

func (c *httpClient) Post (url string, headers http.Header) (*http.Response, error) {
	return c.do(http.MethodPost, url, headers, nil)
}

func Put(c *httpClient) {}

func Patch(c *httpClient) {}

func Delete(c *httpClient) {}