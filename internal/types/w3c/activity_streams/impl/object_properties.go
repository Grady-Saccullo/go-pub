package impl

import "github.com/Grady-Saccullo/activity-pub-go/internal/types/w3c/activity_streams/vocab"

type ObjectProperties struct {
	PropertyName    vocab.PropertyName
	PropertyId      vocab.PropertyId
	PropertySummary vocab.PropertySummary
}

func deserializeObjectProperties(d map[string]interface{}, ldAliases map[string]string, i vocab.ObjectPropertySetters) error {
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

func (o *ObjectProperties) GetPropertyName() vocab.PropertyName {
	return o.PropertyName
}

func (o *ObjectProperties) SetPropertyName(v vocab.PropertyName) {
	o.PropertyName = v
}

func (o *ObjectProperties) GetPropertySummary() vocab.PropertySummary {
	return o.PropertySummary
}

func (o *ObjectProperties) SetPropertySummary(v vocab.PropertySummary) {
	o.PropertySummary = v
}

func (o *ObjectProperties) GetPropertyId() vocab.PropertyId {
	return o.PropertyId
}

func (o *ObjectProperties) SetPropertyId(v vocab.PropertyId) {
	o.PropertyId = v
}
