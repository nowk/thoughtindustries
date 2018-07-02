package endpoint

import (
	"net/url"
	"path"
	"strings"
)

const (
	API_PROTOCOL = "https"
	API_HOST     = "thoughtindustries.com"
	API_VERSION  = "v2"
)

type Endpointer interface {
	Path(...string) string
	Namespace(string)
	Clone() Endpointer
}

type endpoint struct {
	prefix    string
	namespace string
}

func New(prefix string) *endpoint {
	return &endpoint{
		prefix: prefix,
	}
}

// Namespace sets the namespaces for the API repesenting different parts of the
// thougth industries API
func (e *endpoint) Namespace(ns string) {
	e.namespace = ns
}

// Path constructs a final API url
func (e *endpoint) Path(p ...string) string {
	pre := []string{"incoming", API_VERSION}
	if e.namespace != "" {
		pre = append(pre, e.namespace)
	}

	p = append(pre, p...)

	var URL = url.URL{
		Scheme: API_PROTOCOL,
		Host:   strings.Join([]string{e.prefix, API_HOST}, "."),
		Path:   path.Join(p...),
	}

	return URL.String()
}

// Clone returns a new endpoint using the current prefix
// NOTE the namespace is always reset when calling this
func (e *endpoint) Clone() Endpointer {
	return New(e.prefix)
}
