package activity_streams_v2_vocab

import "net/url"

type PropertyActor interface {
	GetIRI() *url.URL

	SetIRI(url.URL)

	IsIRI() bool

	IsObject() bool

	Object
}
