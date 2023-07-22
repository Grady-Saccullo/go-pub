package vocab

import "net/url"

type PropertyTarget interface {
	GetIRI() *url.URL

	SetIRI(url.URL)

	IsIRI() bool

	IsObject() bool

	Object
}
