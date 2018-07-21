package notifications

import (
	"github.com/nowk/thoughtindustries/client"
	. "github.com/nowk/thoughtindustries/data"
)

type CollectionPurchases struct {
	PageInfo PageInfo   `json:"pageInfo"`
	Events   []Purchase `json:"events"`
}

type Collectioner interface {
	Purchases() (*CollectionPurchases, error)
}

type collection struct {
	client.Clienter
}

func (c *collection) Purchases() (*CollectionPurchases, error) {
	var data CollectionPurchases

	err := c.Get("discountGroupPurchase").Json(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
