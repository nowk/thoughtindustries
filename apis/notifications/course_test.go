package notifications

import (
	"net/http"
	"os"
	"testing"

	"github.com/nowk/thoughtindustries/client"
)

type mockReq struct {
	DoFunc func(*http.Request) (*http.Response, error)
}

func (m *mockReq) Do(req *http.Request) (*http.Response, error) {
	return m.DoFunc(req)
}

func TestCoursePurchasesData(t *testing.T) {
	f, err := os.Open("../../data/.json/purchase.json")
	check(t, err)
	mock := &mockReq{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			res := &http.Response{
				Body: f,
			}

			return res, nil
		},
	}

	var cli = client.New("mystuff", "abcdefghijk")
	cli.Req = mock

	evt := New(cli)

	data, err := evt.Course().Purchases()
	if err != nil {
		t.Errorf("expected no error, got %s", err)
	}

	{
		var exp = "MTUxODM2NTg1Mzk5OQ"
		var got = data.PageInfo.Cursor

		if exp != got {
			t.Errorf("expected %s to equal %s", exp, got)
		}
	}

	{
		var exp = "66b131f7-33b7-4d5a-b0fa-bed64acd4c4e-1517868872-ch_x44IEO4smDLAidcy2xz5ZPMg-0"
		var got = data.Events[0].Id

		if exp != got {
			t.Errorf("expected %s to equal %s", exp, got)
		}
	}

	{
		var exp = "1bfaca09-16b3-4f8e-88c3-878e11b0445b"
		var got = data.Events[0].UserDetail.Id

		if exp != got {
			t.Errorf("expected %s to equal %s", exp, got)
		}
	}
}

func check(t *testing.T, err error) {
	if err == nil {
		return
	}

	t.Fatalf("failed to open data test json, %s", err)
}
