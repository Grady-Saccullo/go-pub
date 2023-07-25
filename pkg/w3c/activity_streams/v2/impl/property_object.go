package impl

import (
	"github.com/Grady-Saccullo/go-pub/pkg/w3c/activity_streams/v2/vocab"
	"github.com/Grady-Saccullo/go-pub/pkg/w3c/json_ld/v1/helpers"
	json_ld_v1_impl "github.com/Grady-Saccullo/go-pub/pkg/w3c/json_ld/v1/impl"
	json_ld_v1_vocab "github.com/Grady-Saccullo/go-pub/pkg/w3c/json_ld/v1/vocab"
)

const PropertyObjectKey = "object"

type PropertyObject struct {
	vocab.Object
	vocab.Activity
	json_ld_v1_vocab.TypeIRI
	unknown interface{}
	alias   *string
}

func DeserializePropertyObject(d map[string]interface{}, ldAliases map[string]string) ([]vocab.PropertyObject, error) {
	alias := helpers.GetJsonLDContext(ldAliases, "https://www.w3.org/ns/activitystreams")

	prop, ok := helpers.GetProperty(d, alias, PropertyObjectKey)

	if !ok {
		return nil, nil
	}

	var ret []vocab.PropertyObject

	if list, ok := prop.([]interface{}); ok {
		for _, item := range list {
			if i := deserializePropertyObjectItem(item, ldAliases, alias); i != nil {
				ret = append(ret, i)
			}
		}
	} else {
		if i := deserializePropertyObjectItem(prop, ldAliases, alias); i != nil {
			ret = append(ret, i)
		}
	}

	return ret, nil
}

func deserializePropertyObjectItem(prop interface{}, ldAliases map[string]string, alias *string) *PropertyObject {
	ret := &PropertyObject{
		Object:   &Object{},
		Activity: &Activity{},
		TypeIRI:  &json_ld_v1_impl.TypeIRI{},
		unknown:  nil,
		alias:    alias,
	}

	if v, err := json_ld_v1_impl.DeserializeTypeIRI(prop); err == nil {
		ret.TypeIRI = v
	}

	if j, ok := prop.(map[string]interface{}); ok {
		if o, err := DeserializeObject(j, ldAliases); err == nil {
			ret.Object = o
		}

		if a, err := DeserializeActivity(j, ldAliases); err == nil {
			ret.Activity = a
		}
	}

	return ret
}
