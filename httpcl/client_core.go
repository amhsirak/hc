package httpcl

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"net/http"
	"strings"
)

// for the request body, we need to convert the interface to io.Reader
func (c *httpClient) getRequestBody(contentType string, body interface{}) ([]byte, error) {
	if body == nil {
		return nil, nil
	}
	switch strings.ToLower(contentType) {
	case "application/json":
		return json.Marshal(body)
	
	case "application/xml":
		return xml.Marshal(body)
	
	default:
		return json.Marshal(body)
	}
}

func (c *httpClient) do(method string, url string, headers http.Header, body interface{}) (*http.Response, error) {
	client := http.Client{}

	allHeaders := c.getRequestHeaders(headers)

	requestBody, err := c.getRequestBody(allHeaders.Get("Content-Type"), body)
	if err != nil {
		return nil, errors.New("failed to get request body")
	}

	request, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, errors.New("failed to create a new request")
	}

	request.Header = allHeaders

	return client.Do(request)
}

func (c *httpClient) getRequestHeaders(requestHeaders http.Header) http.Header {
	result := make(http.Header)

	// add default and common headers from the HTTP client instance
	for header, value := range c.Headers {
		if len(value) > 0 {
			result.Set(header, value[0])
		}
	}

	// add custom headers from the request
	// requestHeaders is of type Header map[string][]string
	for header, value := range requestHeaders {
		if len(value) > 0 {
			result.Set(header, value[0])
		}
	}
	return result
}
