package vocab

import "net/url"

type TypeIRI interface {
	GetIRI() *url.URL

	SetIRI(*url.URL)
}
