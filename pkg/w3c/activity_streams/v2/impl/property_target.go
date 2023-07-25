package impl

import (
	"github.com/Grady-Saccullo/go-pub/pkg/w3c/activity_streams/v2/vocab"
	"github.com/Grady-Saccullo/go-pub/pkg/w3c/json_ld/v1/helpers"
	json_ld_v1_impl "github.com/Grady-Saccullo/go-pub/pkg/w3c/json_ld/v1/impl"
	json_ld_v1_vocab "github.com/Grady-Saccullo/go-pub/pkg/w3c/json_ld/v1/vocab"
)

const PropertyTargetKey = "target"

type PropertyTarget struct {
	vocab.Object
	json_ld_v1_vocab.TypeIRI
	unknown interface{}
	alias   *string
}

func DeserializePropertyTarget(d map[string]interface{}, ldAliases map[string]string) ([]vocab.PropertyTarget, error) {
	alias := helpers.GetJsonLDContext(ldAliases, "https://www.w3.org/ns/activitystreams")

	prop, ok := helpers.GetProperty(d, alias, PropertyTargetKey)

	if !ok {
		return nil, nil
	}

	var ret []vocab.PropertyTarget

	if list, ok := prop.([]interface{}); ok {
		for _, item := range list {
			if i := deserializePropertyTargetItem(item, ldAliases, alias); i != nil {
				ret = append(ret, i)
			}
		}
	} else {
		if i := deserializePropertyTargetItem(prop, ldAliases, alias); i != nil {
			ret = append(ret, i)
		}
	}

	return ret, nil
}

func deserializePropertyTargetItem(prop interface{}, ldAliases map[string]string, alias *string) *PropertyTarget {
	ret := &PropertyTarget{
		Object:  &Object{},
		TypeIRI: &json_ld_v1_impl.TypeIRI{},
		unknown: nil,
		alias:   alias,
	}

	if v, err := json_ld_v1_impl.DeserializeTypeIRI(prop); err == nil {
		ret.TypeIRI = v
	}

	if j, ok := prop.(map[string]interface{}); ok {
		if o, err := DeserializeObject(j, ldAliases); err == nil {
			ret.Object = o
		}
	}

	return ret
}
