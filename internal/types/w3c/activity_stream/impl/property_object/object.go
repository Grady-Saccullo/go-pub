package property_object

import (
	"github.com/Grady-Saccullo/activity-pub-go/internal/types/json_ld"
	"github.com/Grady-Saccullo/activity-pub-go/internal/types/w3c/activity_stream/vocab"
	"net/url"
)

const KeyObject = "object"

type PropertyObject struct {
	vocab.ObjectType
	iri     *url.URL
	unknown interface{}
	alias   *string
}

func DeserializePropertyObject(d map[string]interface{}, ldAliases map[string]string) ([]vocab.PropertyObject, error) {
	alias := json_ld.GetJsonLDContext(ldAliases, "https://www.w3.org/ns/activitystreams")

	prop, ok := json_ld.GetProperty(d, alias, KeyObject)

	if !ok {
		return nil, nil
	}

	var ret []vocab.PropertyObject

	if list, ok := prop.([]interface{}); ok {
		for _, item := range list {
			if i := deserializeItem(item, ldAliases); i != nil {
				ret = append(ret, i)
			}
		}
	} else {
		if i := deserializeItem(prop, ldAliases); i != nil {
			ret = append(ret, i)
		}
	}

	return ret, nil
}

func deserializeItem(prop interface{}, ldAliases map[string]string) *PropertyObject {
	if v, ok := json_ld.GetIRI(prop); ok {
		return &PropertyObject{
			iri: v,
		}
	}

	if j, ok := prop.(map[string]interface{}); ok {
		v, err := mgr.DeserializeObjectType()(j, ldAliases)
		if err != nil || v == nil {
			return nil
		}

		return &PropertyObject{
			ObjectType: v,
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
