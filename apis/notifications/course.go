package notifications

import (
	"net/url"

	"github.com/nowk/thoughtindustries/client"
	. "github.com/nowk/thoughtindustries/data"
)

type CoursePurchases struct {
	PageInfo PageInfo   `json:"pageInfo"`
	Events   []Purchase `json:"events"`
}

type CourseViews struct {
	PageInfo PageInfo `json:"pageInfo"`
	Events   []View   `json:"events"`
}

type Courser interface {
	Purchases(...url.Values) (*CoursePurchases, error)
	Views(...url.Values) (*CourseViews, error)
}

type course struct {
	client.Clienter
}

func (c *course) Purchases(qs ...url.Values) (*CoursePurchases, error) {
	var data CoursePurchases

	err := c.Get("coursePurchase", qs...).Json(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (c *course) Views(qs ...url.Values) (*CourseViews, error) {
	var data CourseViews

	err := c.Get("courseView", qs...).Json(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
