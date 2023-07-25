package impl

import (
	"github.com/Grady-Saccullo/go-pub/pkg/w3c/activity_streams/v2/vocab"
	"github.com/Grady-Saccullo/go-pub/pkg/w3c/json_ld/v1/helpers"
	json_ld_v1_impl "github.com/Grady-Saccullo/go-pub/pkg/w3c/json_ld/v1/impl"
	json_ld_v1_vocab "github.com/Grady-Saccullo/go-pub/pkg/w3c/json_ld/v1/vocab"
)

const PropertyIdKey = "id"

type PropertyId struct {
	json_ld_v1_vocab.TypeIRI
	alias *string
}

func DeserializePropertyId(d map[string]interface{}, ldAliases map[string]string) (vocab.PropertyId, error) {
	alias := helpers.GetJsonLDContext(ldAliases, "https://www.w3.org/ns/activitystreams")

	prop, ok := helpers.GetProperty(d, alias, PropertyIdKey)

	if ok {
		v, err := json_ld_v1_impl.DeserializeTypeIRI(prop)
		if err != nil {
			return nil, err
		}
		return &PropertyId{
			TypeIRI: v,
			alias:   alias,
		}, nil
	}

	return &PropertyId{
		TypeIRI: &json_ld_v1_impl.TypeIRI{},
		alias:   alias,
	}, nil
}
