package activity_streams_v2_impl

import (
	"github.com/Grady-Saccullo/activity-pub-go/pkg/w3c/activity_streams/v2/vocab"
)

type ObjectProperties struct {
	PropertyName    activity_streams_v2_vocab.PropertyName
	PropertyId      activity_streams_v2_vocab.PropertyId
	PropertySummary activity_streams_v2_vocab.PropertySummary
}

func deserializeObjectProperties(d map[string]interface{}, ldAliases map[string]string, i activity_streams_v2_vocab.ObjectPropertySetters) error {
	if v, err := DeserializePropertyName(d, ldAliases); err != nil {
		return err
	} else if v != nil {
		i.SetPropertyName(v)
	}

	if v, err := DeserializePropertySummary(d, ldAliases); err != nil {
		return err
	} else if v != nil {
		i.SetPropertySummary(v)
	}

	if v, err := DeserializePropertyId(d, ldAliases); err != nil {
		return err
	} else if v != nil {
		i.SetPropertyId(v)
	}

	return nil
}

func (o *ObjectProperties) GetPropertyName() activity_streams_v2_vocab.PropertyName {
	return o.PropertyName
}

func (o *ObjectProperties) SetPropertyName(v activity_streams_v2_vocab.PropertyName) {
	o.PropertyName = v
}

func (o *ObjectProperties) GetPropertySummary() activity_streams_v2_vocab.PropertySummary {
	return o.PropertySummary
}

func (o *ObjectProperties) SetPropertySummary(v activity_streams_v2_vocab.PropertySummary) {
	o.PropertySummary = v
}

func (o *ObjectProperties) GetPropertyId() activity_streams_v2_vocab.PropertyId {
	return o.PropertyId
}

func (o *ObjectProperties) SetPropertyId(v activity_streams_v2_vocab.PropertyId) {
	o.PropertyId = v
}
