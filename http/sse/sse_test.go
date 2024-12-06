package sse

import (
	"net/http"
	"testing"
)

func TestHandleSSE(t *testing.T) {
	http.HandleFunc("/sse", HandleSSE)
	t.Fatal(http.ListenAndServe(":8080", nil))
}
