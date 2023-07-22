package impl

import (
	"github.com/Grady-Saccullo/activity-pub-go/internal/types/json_ld"
	"github.com/Grady-Saccullo/activity-pub-go/internal/types/w3c/activity_streams/vocab"
	"net/url"
)

const PropertyIdKey = "id"

type PropertyId struct {
	iri   *url.URL
	alias *string
}

func DeserializePropertyId(d map[string]interface{}, ldAliases map[string]string) (vocab.PropertyId, error) {
	alias := json_ld.GetJsonLDContext(ldAliases, "https://www.w3.org/ns/activitystreams")

	prop, ok := json_ld.GetProperty(d, alias, PropertyIdKey)

	if ok {
		if v, ok := json_ld.GetIRI(prop); ok {
			return &PropertyId{
				iri:   v,
				alias: alias,
			}, nil
		}
	}

	return nil, nil
}

func (p *PropertyId) GetIRI() *url.URL {
	return p.iri
}

func (p *PropertyId) SetIRI(iri url.URL) {
	p.iri = &iri
}

func (p *PropertyId) IsIRI() bool {
	return p.iri != nil
}
