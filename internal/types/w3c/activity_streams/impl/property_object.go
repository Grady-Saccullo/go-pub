package impl

import (
	"github.com/Grady-Saccullo/activity-pub-go/internal/types/json_ld"
	"github.com/Grady-Saccullo/activity-pub-go/internal/types/w3c/activity_streams/vocab"
	"net/url"
)

const PropertyObjectKey = "object"

type PropertyObject struct {
	vocab.Object
	vocab.Activity
	iri     *url.URL
	unknown interface{}
	alias   *string
}

func DeserializePropertyObject(d map[string]interface{}, ldAliases map[string]string) ([]vocab.PropertyObject, error) {
	alias := json_ld.GetJsonLDContext(ldAliases, "https://www.w3.org/ns/activitystreams")

	prop, ok := json_ld.GetProperty(d, alias, PropertyObjectKey)

	if !ok {
		return nil, nil
	}

	var ret []vocab.PropertyObject

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
	if v, ok := json_ld.GetIRI(prop); ok {
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
