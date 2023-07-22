package impl

import (
	"github.com/Grady-Saccullo/activity-pub-go/internal/types/json_ld"
	"github.com/Grady-Saccullo/activity-pub-go/internal/types/w3c/activity_streams/vocab"
	"net/url"
)

const PropertyTargetKey = "target"

type PropertyTarget struct {
	vocab.Object
	iri     *url.URL
	unknown interface{}
	alias   *string
}

func DeserializePropertyTarget(d map[string]interface{}, ldAliases map[string]string) ([]vocab.PropertyTarget, error) {
	alias := json_ld.GetJsonLDContext(ldAliases, "https://www.w3.org/ns/activitystreams")

	prop, ok := json_ld.GetProperty(d, alias, PropertyTargetKey)

	if !ok {
		return nil, nil
	}

	var ret []vocab.PropertyTarget

	if list, ok := prop.([]interface{}); ok {
		for _, item := range list {
			if i := deserializePropertyTargetItem(item, ldAliases); i != nil {
				ret = append(ret, i)
			}
		}
	} else {
		if i := deserializePropertyTargetItem(prop, ldAliases); i != nil {
			ret = append(ret, i)
		}
	}

	return ret, nil
}

func deserializePropertyTargetItem(prop interface{}, ldAliases map[string]string) *PropertyTarget {
	if v, ok := json_ld.GetIRI(prop); ok {
		return &PropertyTarget{
			iri: v,
		}
	}

	if j, ok := prop.(map[string]interface{}); ok {
		v, err := DeserializeObject(j, ldAliases)
		if err != nil || v == nil {
			return nil
		}

		return &PropertyTarget{
			Object: v,
		}
	}

	return nil
}

func (p *PropertyTarget) GetIRI() *url.URL {
	return p.iri
}

func (p *PropertyTarget) SetIRI(iri url.URL) {
	p.iri = &iri
}

func (p *PropertyTarget) IsIRI() bool {
	return p.iri != nil
}

func (p *PropertyTarget) IsObject() bool {
	return p.Object != nil
}
