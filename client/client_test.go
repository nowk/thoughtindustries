package client

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"
)

func TestCloneReturnsAClientWithAClonedEndpoint(t *testing.T) {
	a := New("abcd", "1234")

	b := a.Clone()
	b.Endpoint().Namespace("foo")

	{
		pathb := b.Endpoint().Path("bar")

		var exp = "https://abcd.thoughtindustries.com/incoming/v2/foo/bar"
		var got = pathb

		if exp != got {
			t.Errorf("expected %s, got %s", exp, got)
		}
	}

	{
		patha := a.Endpoint().Path("bar")

		var exp = "https://abcd.thoughtindustries.com/incoming/v2/bar"
		var got = patha

		if exp != got {
			t.Errorf("expected %s, got %s", exp, got)
		}
	}
}

type mockReq struct {
	DoFunc func(*http.Request) (*http.Response, error)
}

func (m *mockReq) Do(req *http.Request) (*http.Response, error) {
	return m.DoFunc(req)
}

func TestPing(t *testing.T) {
	body := strings.NewReader(`{"valid": true}`)
	mock := &mockReq{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			res := &http.Response{
				Body: ioutil.NopCloser(body),
			}

			return res, nil
		},
	}

	var cli = New("mystuff", "abcdefghijk")
	cli.Req = mock

	status, err := cli.Ping()
	if err != nil {
		t.Errorf("expected no error, got %s", err)
	}

	var exp = true
	var got = status.Valid

	if exp != got {
		t.Errorf("expected status to be true")
	}
}

func TestPingLive(t *testing.T) {
	if os.Getenv("TEST_LIVE") != "true" {
		t.Skipf("To run Live times please set env TEST_LIVE=true")
	}

	var (
		l = os.Getenv("THOUGHTINDUSTRIES_LOGIN")
		k = os.Getenv("THOUGHTINDUSTRIES_API_KEY")

		cli = New(l, k)
	)
	status, err := cli.Ping()
	if err != nil {
		t.Errorf("expected no error, got %s", err)
	}

	t.Logf("status: %v", status)
}
