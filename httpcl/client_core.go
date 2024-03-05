package httpcl

import (
	"errors"
	"net/http"
)

func (c *httpClient) do(method string, url string, headers http.Header, body interface{}) (*http.Response, error) {
	client := http.Client{}

	request, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, errors.New("failed to create a new request")
	}

	return client.Do(request)
}

func (c *httpClient) getRequestHeaders() http.Header {
	result := make(http.Header)

	// Default | Common Headers For Request 
	for header, value := range c.Headers {
		if len(value) > 0 {
			request.Header.Set(header, value[0])
		}
	}

	// Custom Headers For Request
	// headers is of type Header map[string][]string
	for header, value := range headers {
		if len(value) > 0 {
			request.Header.Set(header, value[0])
		}
	}
	return result
}
