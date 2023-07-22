package activity_streams_v2_impl

import (
	"github.com/Grady-Saccullo/activity-pub-go/pkg/w3c/activity_streams/v2/vocab"
	"github.com/Grady-Saccullo/activity-pub-go/pkg/w3c/json_ld/v1"
	"net/url"
)

const PropertyActorKey = "actor"

type PropertyActor struct {
	activity_streams_v2_vocab.Object
	iri     *url.URL
	unknown interface{}
	alias   *string
}

func DeserializePropertyActor(d map[string]interface{}, ldAliases map[string]string) ([]activity_streams_v2_vocab.PropertyActor, error) {
	alias := json_ld_v1.GetJsonLDContext(ldAliases, "https://www.w3.org/ns/activitystreams")

	prop, ok := json_ld_v1.GetProperty(d, alias, PropertyActorKey)

	if !ok {
		return nil, nil
	}

	var ret []activity_streams_v2_vocab.PropertyActor

	if list, ok := prop.([]interface{}); ok {
		for _, item := range list {
			if i := deserializePropertyActorItem(item, ldAliases); i != nil {
				ret = append(ret, i)
			}
		}
	} else {
		if i := deserializePropertyActorItem(prop, ldAliases); i != nil {
			ret = append(ret, i)
		}
	}

	return ret, nil
}

func deserializePropertyActorItem(prop interface{}, ldAliases map[string]string) *PropertyActor {
	if v, ok := json_ld_v1.GetIRI(prop); ok {
		return &PropertyActor{
			iri: v,
		}
	}

	if j, ok := prop.(map[string]interface{}); ok {
		v, err := DeserializeObject(j, ldAliases)
		if err != nil || v == nil {
			return nil
		}

		return &PropertyActor{
			Object: v,
		}
	}

	return nil
}

func (p *PropertyActor) GetIRI() *url.URL {
	return p.iri
}

func (p *PropertyActor) SetIRI(iri url.URL) {
	p.iri = &iri
}

func (p *PropertyActor) IsIRI() bool {
	return p.iri != nil
}

func (p *PropertyActor) IsObject() bool {
	return p.Object != nil
}