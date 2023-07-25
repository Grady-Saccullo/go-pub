package impl

import (
	"fmt"
	"github.com/Grady-Saccullo/go-pub/pkg/w3c/activity_streams/v2/vocab"
	"github.com/Grady-Saccullo/go-pub/pkg/w3c/json_ld/v1/helpers"
	rdf_v1 "github.com/Grady-Saccullo/go-pub/pkg/w3c/rdf/v1"
)

const PropertyNameKey = "name"
const PropertyNameKeyMap = "nameMap"

type PropertyName struct {
	rdfLangStringMap map[string]string
	str              *string
	alias            *string
}

func DeserializePropertyName(d map[string]interface{}, ldAliases map[string]string) (vocab.PropertyName, error) {
	alias := helpers.GetJsonLDContext(ldAliases, "https://www.w3.org/ns/activitystreams")

	propSingle, okSingle := helpers.GetProperty(d, alias, PropertyNameKey)
	propMap, okMap := helpers.GetProperty(d, alias, PropertyNameKeyMap)

	if okSingle && okMap {
		return nil, fmt.Errorf("both %s and %s keys are present", PropertyNameKey, PropertyNameKeyMap)
	}

	ret := &PropertyName{
		alias: alias,
	}

	if okSingle {
		if v, ok := propSingle.(string); ok {
			ret.str = &v
		}
	} else if okMap {
		if v, err := rdf_v1.DeserializeLangString(propMap); err != nil {
			return nil, err
		} else {
			ret.rdfLangStringMap = v
		}
	}

	return ret, nil
}

func (p *PropertyName) GetString() *string {
	return p.str
}

func (p *PropertyName) SetString(str string) {
	p.str = &str
}

func (p *PropertyName) GetRDFLangStringMap() *map[string]string {
	return &p.rdfLangStringMap
}

func (p *PropertyName) SetRDFLangStringMap(strMap map[string]string) {
	p.rdfLangStringMap = strMap
}
