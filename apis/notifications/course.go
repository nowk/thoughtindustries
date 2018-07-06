package notifications

import (
	"github.com/nowk/thoughtindustries/client"
	. "github.com/nowk/thoughtindustries/data"
)

type CoursePurchases struct {
	PageInfo PageInfo   `json:"pageInfo"`
	Events   []Purchase `json:"events"`
}

type CourseViews struct {
	PageInfo PageInfo `json:"pageInfo"`
	Views    []View   `json:"events"`
}

type Courser interface {
	Purchases() (*CoursePurchases, error)
	Views() (*CourseViews, error)
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

func (c *course) Views() (*CourseViews, error) {
	var data CourseViews

	err := c.Get("courseView").Json(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
