package impl

import (
	"github.com/Grady-Saccullo/go-pub/pkg/w3c/json_ld/v1/vocab"
	"net/url"
)

type TypeIRI struct {
	iri *url.URL
}

func DeserializeTypeIRI(p interface{}) (vocab.TypeIRI, error) {
	if v, ok := p.(string); ok {
		if u, err := url.Parse(v); err != nil {
			return nil, err
		} else {
			return &TypeIRI{
				iri: u,
			}, nil
		}
	}

	return &TypeIRI{
		iri: nil,
	}, nil
}

func (t *TypeIRI) GetIRI() *url.URL {
	return t.iri
}

func (t *TypeIRI) SetIRI(u *url.URL) {
	t.iri = u
}
