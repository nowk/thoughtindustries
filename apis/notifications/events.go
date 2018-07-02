package notifications

import (
	"github.com/nowk/thoughtindustries/client"
)

type Eventser interface {
	Course() Courser
}

type events struct {
	course Courser
}

func New(c client.Clienter) Eventser {
	cl := c.Clone()
	cl.Endpoint().Namespace("events")

	return &events{
		course: &course{cl},
	}
}

func (e *events) Course() Courser {
	return e.course
}
