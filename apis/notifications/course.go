package notifications

import (
	"github.com/nowk/thoughtindustries/client"
	. "github.com/nowk/thoughtindustries/data"
)

type CoursePurchases struct {
	PageInfo PageInfo   `json:"pageInfo"`
	Events   []Purchase `json:"events"`
}

type Courser interface {
	Purchases() (*CoursePurchases, error)
}

type course struct {
	client.Clienter
}

func (c *course) Purchases() (*CoursePurchases, error) {
	var data CoursePurchases

	err := c.Get("coursePurchase").Json(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
