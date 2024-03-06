package httpcl

import (
	"net/http"
	"testing"
)

// flow of tests: initalization, execution, validation

func TestGetRequestHeaders(t *testing.T) {
	client := httpClient{}
	commonHeaders := make(http.Header)
	commonHeaders.Set("Content-Type", "application/json")
	commonHeaders.Set("User-Agent", "test-httpcl")
	client.Headers = commonHeaders

	requestHeaders := make(http.Header)
	requestHeaders.Set("X-Request-Id", "ABC-999")
	finalHeaders := client.getRequestHeaders(requestHeaders)

	if len (finalHeaders) != 3 {
		t.Error("Expected 3 headers, got ", len(finalHeaders))
	}
}