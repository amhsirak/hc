package httpcl

import (
	"net/http"
	"testing"
)

// flow of tests: initalization*, execution, validation

func TestGetRequestHeaders(t *testing.T) {
	client := httpClient{}
	commonHeaders := make(http.Header)
	commonHeaders.Set("Content-Type", "application/json")
	commonHeaders.Set("User-Agent", "test-httpcl")
	client.Headers = commonHeaders

	requestHeaders := make(http.Header)
	requestHeaders.Set("X-Request-Id", "ABC-999")
	finalHeaders := client.getRequestHeaders(requestHeaders)

	if len(finalHeaders) != 3 {
		t.Error("expected 3 headers, got ", len(finalHeaders))
	}

	if finalHeaders.Get("X-Request-Id") != "ABC-999" {
		t.Error("invalid request id, expected 'ABC-999', got ", finalHeaders.Get("X-Request-Id"))
	}

	if finalHeaders.Get("Content-Type") != "application/json" {
		t.Error("invalid content type, expected 'application/json', got ", finalHeaders.Get("Content-Type"))
	}

	if finalHeaders.Get("User-Agent") != "test-httpcl" {
		t.Error("invalid user agent type, expected 'test-httpcl', got ", finalHeaders.Get("Content-Type"))
	}
}

func TestGetRequestBody(t *testing.T) {
	client := httpClient{}

	t.Run("NoBodyNilResponse", func(t *testing.T) {
		
		body, err := client.getRequestBody("", nil)

		if err != nil {
			t.Error("no error expected when passing a nil body")
		}

		if body != nil {
			t.Error("no body expected when passing a nil body")
		}
	})

	t.Run("BodyWithJson", func(t *testing.T) {
		requestBody := []string{"one", "two"}

		body, err := client.getRequestBody("application/json", requestBody)

		if err != nil {
			t.Error("no error expected when marshaling slice as json")
		}

		if string(body) != `["one","two"]` {
			t.Error("invalid json body obtained")
		}
	})

	t.Run("BodyWithXml", func(t *testing.T) {
		requestBody := []string{"one", "two"}

		body, err := client.getRequestBody("application/xml", requestBody)

		if err != nil {
			t.Error("no error expected when marshaling slice as xml")
		}

		if string(body) != `<string>one</string><string>two</string>` {
			t.Error("invalid xml body obtained")
		}
	})

}
