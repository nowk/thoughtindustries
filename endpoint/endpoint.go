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
	Path([]url.Values, ...string) string
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
func (e *endpoint) Path(qs []url.Values, p ...string) string {
	pre := []string{"incoming", API_VERSION}
	if e.namespace != "" {
		pre = append(pre, e.namespace)
	}

	p = append(pre, p...)

	var URL = url.URL{
		Scheme:   API_PROTOCOL,
		Host:     strings.Join([]string{e.prefix, API_HOST}, "."),
		Path:     path.Join(p...),
		RawQuery: queryString(qs),
	}

	return URL.String()
}

// queryString returns the raw query for a given set of url.Values
func queryString(qs []url.Values) string {
	if len(qs) == 0 {
		return ""
	}

	q := url.Values{}
	for _, n := range qs {
		for k, _ := range n {
			q.Set(k, n.Get(k))
		}
	}

	return q.Encode()
}

// Clone returns a new endpoint using the current prefix
// NOTE the namespace is always reset when calling this
func (e *endpoint) Clone() Endpointer {
	return New(e.prefix)
}
