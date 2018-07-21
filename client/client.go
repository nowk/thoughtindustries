package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/nowk/thoughtindustries/endpoint"
)

const (
	GET = "GET"
)

// httpClienter is a basic http client interface
type httpClienter interface {
	Do(*http.Request) (*http.Response, error)
}

// Decoder is an interface to wrap encoding calls
type Decoder interface {
	Json(interface{}) error
}

type decode struct {
	io.ReadCloser

	// err holds any previous errors from the chain and will be returned on the
	// next method call if it is not nil
	err error
}

// Json decodes to JSON, this will close the ReadCloser
func (d *decode) Json(v interface{}) error {
	// we have an error set return it
	if d.err != nil {
		return d.err
	}

	defer d.Close()

	return json.NewDecoder(d).Decode(v)
}

type Clienter interface {
	Get(string, ...url.Values) Decoder
	Clone() Clienter
	Endpoint() endpoint.Endpointer
}

type client struct {
	Req      httpClienter        // Req is exported to allow configuration
	endpoint endpoint.Endpointer //

	token string
}

func New(prefix, token string) *client {
	return &client{
		Req:      http.DefaultClient,
		endpoint: endpoint.New(prefix),

		token: token,
	}
}

// Get makes the actual http request and returns a Decoder interface to decode
// to an object of your choice
func (c *client) Get(path string, qs ...url.Values) Decoder {
	var url = c.Endpoint().Path(qs, path)

	req, err := http.NewRequest(GET, url, nil)
	if err != nil {
		return &decode{err: err}
	}

	c.authorize(req)

	res, err := c.Req.Do(req)
	if err != nil {
		return &decode{err: err}
	}

	return &decode{ReadCloser: res.Body}
}

// Clone returns a new client using the current httpClienter and token with a
// cloned endpointer
func (c *client) Clone() Clienter {
	return &client{
		Req:      c.Req,
		endpoint: c.Endpoint().Clone(),

		token: c.token,
	}
}

func (c *client) Endpoint() endpoint.Endpointer {
	return c.endpoint
}

type Status struct {
	Valid bool `json:"valid"`
}

// Pint calls the thoughtindustries /ping endpoint
// NOTE this is not part of the Clienter interface, it's is only available to
// the original client instance created when calling New
func (c *client) Ping() (*Status, error) {
	var (
		status Status

		err = c.Get("ping").Json(&status)
	)
	if err != nil {
		return nil, err
	}

	return &status, nil
}

func (c *client) bearer() string {
	return fmt.Sprintf("Bearer %s", c.token)
}

func (c *client) authorize(req *http.Request) *http.Request {
	req.Header.Add("Authorization", c.bearer())

	return req
}
