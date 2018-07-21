package notifications

import (
	"net/url"

	"github.com/nowk/thoughtindustries/client"
	. "github.com/nowk/thoughtindustries/data"
)

type CollectionPurchases struct {
	PageInfo PageInfo   `json:"pageInfo"`
	Events   []Purchase `json:"events"`
}

type Collectioner interface {
	Purchases(...url.Values) (*CollectionPurchases, error)
}

type collection struct {
	client.Clienter
}

func (c *collection) Purchases(qs ...url.Values) (*CollectionPurchases, error) {
	var data CollectionPurchases

	err := c.Get("discountGroupPurchase", qs...).Json(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
