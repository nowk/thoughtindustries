package notifications

import (
	"net/http"
	"testing"
)

// keeping some basic structures used to test here

type mockReq struct {
	DoFunc func(*http.Request) (*http.Response, error)
}

func (m *mockReq) Do(req *http.Request) (*http.Response, error) {
	return m.DoFunc(req)
}

func check(t *testing.T, err error) {
	if err == nil {
		return
	}

	t.Fatalf("failed to open data test json, %s", err)
}
