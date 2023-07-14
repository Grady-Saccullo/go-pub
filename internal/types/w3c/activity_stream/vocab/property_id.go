package vocab

import "net/url"

type PropertyId interface {
	GetIRI() *url.URL

	SetIRI(url.URL)

	IsIRI() bool
}
