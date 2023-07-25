package impl

import (
	"fmt"
	"github.com/Grady-Saccullo/go-pub/pkg/w3c/activity_streams/v2/vocab"
	"github.com/Grady-Saccullo/go-pub/pkg/w3c/json_ld/v1/helpers"
	rdf_v1 "github.com/Grady-Saccullo/go-pub/pkg/w3c/rdf/v1"
)

const KeySummary = "summary"
const KeySummaryMap = "summaryMap"

type PropertySummary struct {
	unknown          interface{}
	rdfLangStringMap map[string]string
	str              string
	alias            *string
}

func DeserializePropertySummary(d map[string]interface{}, ldAliases map[string]string) (vocab.PropertySummary, error) {
	alias := helpers.GetJsonLDContext(ldAliases, "https://www.w3.org/ns/activitystreams")

	propSingle, okSingle := helpers.GetProperty(d, alias, KeySummary)
	propMap, okMap := helpers.GetProperty(d, alias, KeySummaryMap)

	if okSingle && okMap {
		return nil, fmt.Errorf("both %s and %s keys are present", KeySummary, KeySummaryMap)
	}

	ret := &PropertySummary{
		alias: alias,
	}

	if okSingle {
		if v, ok := propSingle.(string); ok {
			ret.str = v
		} else {
			ret.unknown = propSingle
		}
	} else {
		v, err := rdf_v1.DeserializeLangString(propMap)

		if err != nil {
			ret.unknown = propMap
		} else {
			ret.rdfLangStringMap = v
		}
	}

	return ret, nil
}

func (p *PropertySummary) GetString() *string {
	return &p.str
}

func (p *PropertySummary) SetString(str string) {
	p.str = str
}

func (p *PropertySummary) GetStringMap() *map[string]string {
	return &p.rdfLangStringMap
}

func (p *PropertySummary) SetStringMap(strMap map[string]string) {
	p.rdfLangStringMap = strMap
}

func (p *PropertySummary) GetUnknown() interface{} {
	return p.unknown
}
