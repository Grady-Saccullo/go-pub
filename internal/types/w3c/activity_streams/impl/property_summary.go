package impl

import (
	"fmt"
	"github.com/Grady-Saccullo/activity-pub-go/internal/types/json_ld"
	"github.com/Grady-Saccullo/activity-pub-go/internal/types/rdf"
	"github.com/Grady-Saccullo/activity-pub-go/internal/types/w3c/activity_streams/vocab"
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
	alias := json_ld.GetJsonLDContext(ldAliases, "https://www.w3.org/ns/activitystreams")

	propSingle, okSingle := json_ld.GetProperty(d, alias, KeySummary)
	propMap, okMap := json_ld.GetProperty(d, alias, KeySummaryMap)

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
		v, err := rdf.DeserializeLangString(propMap)

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
