package impl

import (
	"fmt"
	"github.com/Grady-Saccullo/activity-pub-go/internal/types/json_ld"
	"github.com/Grady-Saccullo/activity-pub-go/internal/types/rdf"
	"github.com/Grady-Saccullo/activity-pub-go/internal/types/w3c/activity_stream/vocab"
)

const PropertyNameKey = "name"
const PropertyNameKeyMap = "nameMap"

type PropertyName struct {
	unknown          interface{}
	rdfLangStringMap map[string]string
	str              string
	alias            *string
}

func DeserializePropertyName(d map[string]interface{}, ldAliases map[string]string) (vocab.PropertyName, error) {
	alias := json_ld.GetJsonLDContext(ldAliases, "https://www.w3.org/ns/activitystreams")

	propSingle, okSingle := json_ld.GetProperty(d, alias, PropertyNameKey)
	propMap, okMap := json_ld.GetProperty(d, alias, PropertyNameKeyMap)

	if okSingle && okMap {
		return nil, fmt.Errorf("both %s and %s keys are present", PropertyNameKey, PropertyNameKeyMap)
	}

	if okSingle {
		if v, ok := propSingle.(string); ok {
			ret := &PropertyName{
				str:   v,
				alias: alias,
			}
			return ret, nil
		} else {
			ret := &PropertyName{
				unknown: propSingle,
				alias:   alias,
			}
			return ret, nil
		}
	}

	if okMap {
		v, err := rdf.DeserializeLangString(propMap)

		if err != nil {
			ret := &PropertyName{
				unknown: propMap,
				alias:   alias,
			}
			return ret, nil
		}

		ret := &PropertyName{
			rdfLangStringMap: v,
			alias:            alias,
		}
		return ret, nil
	}

	return nil, nil
}

func (p *PropertyName) GetString() *string {
	return &p.str
}

func (p *PropertyName) SetString(str string) {
	p.str = str
}

func (p *PropertyName) IsString() bool {
	return len(p.str) != 0
}

func (p *PropertyName) GetRDFLangStringMap() *map[string]string {
	return &p.rdfLangStringMap
}

func (p *PropertyName) SetRDFLangStringMap(strMap map[string]string) {
	p.rdfLangStringMap = strMap
}

func (p *PropertyName) IsRDFLangStringMap() bool {
	return p.rdfLangStringMap != nil
}

func (p *PropertyName) GetUnknown() interface{} {
	return p.unknown
}

func (p *PropertyName) IsUnknown() bool {
	return p.unknown != nil
}
