package internal

import (
	"bytes"
	"net/http"
)

func newRequest(url, body string) *http.Request {
	request, _ := http.NewRequest(http.MethodPost, url, bytes.NewBufferString(body))
	request.Header.Set("Content-Type", "application/json")
	return request
}
