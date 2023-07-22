package activity_streams_v2_impl

import (
	"github.com/Grady-Saccullo/activity-pub-go/pkg/w3c/activity_streams/v2/vocab"
	json_ld_v1 "github.com/Grady-Saccullo/activity-pub-go/pkg/w3c/json_ld/v1"
	"net/url"
)

const PropertyObjectKey = "object"

type PropertyObject struct {
	activity_streams_v2_vocab.Object
	activity_streams_v2_vocab.Activity
	iri     *url.URL
	unknown interface{}
	alias   *string
}

func DeserializePropertyObject(d map[string]interface{}, ldAliases map[string]string) ([]activity_streams_v2_vocab.PropertyObject, error) {
	alias := json_ld_v1.GetJsonLDContext(ldAliases, "https://www.w3.org/ns/activitystreams")

	prop, ok := json_ld_v1.GetProperty(d, alias, PropertyObjectKey)

	if !ok {
		return nil, nil
	}

	var ret []activity_streams_v2_vocab.PropertyObject

	if list, ok := prop.([]interface{}); ok {
		for _, item := range list {
			if i := deserializePropertyObjectItem(item, ldAliases); i != nil {
				ret = append(ret, i)
			}
		}
	} else {
		if i := deserializePropertyObjectItem(prop, ldAliases); i != nil {
			ret = append(ret, i)
		}
	}

	return ret, nil
}

func deserializePropertyObjectItem(prop interface{}, ldAliases map[string]string) *PropertyObject {
	if v, ok := json_ld_v1.GetIRI(prop); ok {
		return &PropertyObject{
			iri: v,
		}
	}

	if j, ok := prop.(map[string]interface{}); ok {
		o, err := DeserializeObject(j, ldAliases)
		if err == nil && o != nil {
			return &PropertyObject{
				iri:    nil,
				Object: o,
			}
		}

		a, err := DeserializeActivity(j, ldAliases)
		if err == nil && a != nil {
			return &PropertyObject{
				iri:      nil,
				Activity: a,
			}
		}

	}

	return nil
}

func (p *PropertyObject) GetIRI() *url.URL {
	return p.iri
}

func (p *PropertyObject) SetIRI(iri url.URL) {
	p.iri = &iri
}

func (p *PropertyObject) IsIRI() bool {
	return p.iri != nil
}

func (p *PropertyObject) IsObject() bool {
	return p.Object != nil
}

func (p *PropertyObject) IsActivity() bool {
	return p.Activity != nil
}
