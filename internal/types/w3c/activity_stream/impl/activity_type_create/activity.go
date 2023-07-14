package activity_type_create

import (
	"fmt"
	"github.com/Grady-Saccullo/activity-pub-go/internal/types/json_ld"
	"github.com/Grady-Saccullo/activity-pub-go/internal/types/w3c/activity_stream/vocab"
)

const TypeValue = "Create"

type ActivityTypeCreate struct {
	alias           *string
	PropertyName    vocab.PropertyName
	PropertySummary vocab.PropertySummary
	PropertyObject  []vocab.PropertyObject
}

func DeserializeActivityTypeCreate(d map[string]interface{}, ldAliases map[string]string) (vocab.ActivityTypeCreate, error) {
	alias := json_ld.GetJsonLDContext(ldAliases, "https://www.w3.org/ns/activitystreams")

	s, ok := json_ld.GetType(d, alias)

	if !ok {
		return nil, fmt.Errorf("type is not defined")
	} else if *s != TypeValue {
		return nil, nil
	}

	ret := ActivityTypeCreate{}

	if v, err := mgr.DeserializePropertyName()(d, ldAliases); err != nil {
		return nil, err
	} else if v != nil {
		ret.PropertyName = v
	}

	if v, err := mgr.DeserializePropertySummary()(d, ldAliases); err != nil {
		return nil, err
	} else if v != nil {
		ret.PropertySummary = v
	}

	if v, err := mgr.DeserializePropertyObject()(d, ldAliases); err != nil {
		return nil, err
	} else if v != nil {
		ret.PropertyObject = v
	}

	return &ret, nil
}

func (n *ActivityTypeCreate) GetPropertyName() vocab.PropertyName {
	return n.PropertyName
}

func (n *ActivityTypeCreate) GetPropertySummary() vocab.PropertySummary {
	return n.PropertySummary
}

func (n *ActivityTypeCreate) GetPropertyObject() []vocab.PropertyObject {
	return n.PropertyObject
}
