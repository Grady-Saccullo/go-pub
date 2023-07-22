package activity_streams_v2_impl

import (
	"github.com/Grady-Saccullo/activity-pub-go/pkg/w3c/activity_streams/v2/vocab"
	"github.com/Grady-Saccullo/activity-pub-go/pkg/w3c/json_ld/v1"
	"net/url"
)

const PropertyIdKey = "id"

type PropertyId struct {
	iri   *url.URL
	alias *string
}

func DeserializePropertyId(d map[string]interface{}, ldAliases map[string]string) (activity_streams_v2_vocab.PropertyId, error) {
	alias := json_ld_v1.GetJsonLDContext(ldAliases, "https://www.w3.org/ns/activitystreams")

	prop, ok := json_ld_v1.GetProperty(d, alias, PropertyIdKey)

	if ok {
		if v, ok := json_ld_v1.GetIRI(prop); ok {
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
