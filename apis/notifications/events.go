package notifications

import (
	"github.com/nowk/thoughtindustries/client"
)

type Eventser interface {
	Course() Courser
	Collection() Collectioner
}

type events struct {
	course     Courser
	collection Collectioner
}

func New(c client.Clienter) Eventser {
	cl := c.Clone()
	cl.Endpoint().Namespace("events")

	return &events{
		course:     &course{cl},
		collection: &collection{cl},
	}
}

func (e *events) Course() Courser {
	return e.course
}

func (e *events) Collection() Collectioner {
	return e.collection
}
