package notifications

import (
	"net/http"
	"os"
	"testing"

	"github.com/nowk/thoughtindustries/client"
)

func TestCollectionPurchasesData(t *testing.T) {
	f, err := os.Open("../../data/.json/collection/purchases.json")
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

	data, err := evt.Collection().Purchases()
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
		var exp = "087cc8ea-fc25-40c9-a433-9c64bea3e475-1516480087-ch_x45fwds16iI23A1aSGz9egeJ-0"
		var got = data.Events[0].Id

		if exp != got {
			t.Errorf("expected %s to equal %s", exp, got)
		}
	}

	{
		var exp = "discountGroup"
		var got = data.Events[0].PurchasableType

		if exp != got {
			t.Errorf("expected %s to equal %s", exp, got)
		}
	}

	{
		var exp = "c1eb0b65-68fe-4cf7-8445-6da329d75e48"
		var got = data.Events[0].UserDetail.Id

		if exp != got {
			t.Errorf("expected %s to equal %s", exp, got)
		}
	}
}
