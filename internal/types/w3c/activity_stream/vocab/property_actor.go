package vocab

import "net/url"

type PropertyActor interface {
	GetIRI() *url.URL

	SetIRI(url.URL)

	IsIRI() bool

	IsObject() bool

	Object
}
