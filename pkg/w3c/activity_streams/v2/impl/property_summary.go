package activity_streams_v2_impl

import (
	"fmt"
	"github.com/Grady-Saccullo/activity-pub-go/pkg/w3c/activity_streams/v2/vocab"
	json_ld_v1 "github.com/Grady-Saccullo/activity-pub-go/pkg/w3c/json_ld/v1"
	v12 "github.com/Grady-Saccullo/activity-pub-go/pkg/w3c/rdf/v1"
)

const KeySummary = "summary"
const KeySummaryMap = "summaryMap"

type PropertySummary struct {
	unknown          interface{}
	rdfLangStringMap map[string]string
	str              string
	alias            *string
}

func DeserializePropertySummary(d map[string]interface{}, ldAliases map[string]string) (activity_streams_v2_vocab.PropertySummary, error) {
	alias := json_ld_v1.GetJsonLDContext(ldAliases, "https://www.w3.org/ns/activitystreams")

	propSingle, okSingle := json_ld_v1.GetProperty(d, alias, KeySummary)
	propMap, okMap := json_ld_v1.GetProperty(d, alias, KeySummaryMap)

	if okSingle && okMap {
		return nil, fmt.Errorf("both %s and %s keys are present", KeySummary, KeySummaryMap)
	}

	if okSingle {
		if v, ok := propSingle.(string); ok {
			ret := &PropertySummary{
				str:   v,
				alias: alias,
			}
			return ret, nil
		} else {
			ret := &PropertySummary{
				unknown: propSingle,
				alias:   alias,
			}
			return ret, nil
		}
	}

	if okMap {
		v, err := v12.DeserializeLangString(propMap)

		if err != nil {
			ret := &PropertySummary{
				unknown: propMap,
				alias:   alias,
			}
			return ret, nil
		}
		ret := &PropertySummary{
			rdfLangStringMap: v,
			alias:            alias,
		}
		return ret, nil
	}

	return nil, nil
}

func (p *PropertySummary) GetString() *string {
	return &p.str
}

func (p *PropertySummary) SetString(str string) {
	p.str = str
}

func (p *PropertySummary) IsString() bool {
	return len(p.str) != 0
}

func (p *PropertySummary) GetStringMap() *map[string]string {
	return &p.rdfLangStringMap
}

func (p *PropertySummary) SetStringMap(strMap map[string]string) {
	p.rdfLangStringMap = strMap
}

func (p *PropertySummary) IsRDFLangStringMap() bool {
	return p.rdfLangStringMap != nil
}

func (p *PropertySummary) GetUnknown() interface{} {
	return p.unknown
}

func (p *PropertySummary) IsUnknown() bool {
	return p.unknown != nil
}
