package impl

import "github.com/Grady-Saccullo/go-pub/pkg/w3c/activity_streams/v2/vocab"

type ObjectProperties struct {
	vocab.PropertyName
	vocab.PropertyId
	vocab.PropertySummary
}

func deserializeObjectProperties(d map[string]interface{}, ldAliases map[string]string, i vocab.ObjectPropertySetters) error {
	if v, err := DeserializePropertyName(d, ldAliases); err != nil {
		return err
	} else if v != nil {
		i.SetName(v)
	}

	if v, err := DeserializePropertySummary(d, ldAliases); err != nil {
		return err
	} else if v != nil {
		i.SetSummary(v)
	}

	if v, err := DeserializePropertyId(d, ldAliases); err != nil {
		return err
	} else if v != nil {
		i.SetId(v)
	}

	return nil
}

func (o *ObjectProperties) GetName() vocab.PropertyName {
	return o.PropertyName
}

func (o *ObjectProperties) SetName(v vocab.PropertyName) {
	o.PropertyName = v
}

func (o *ObjectProperties) GetSummary() vocab.PropertySummary {
	return o.PropertySummary
}

func (o *ObjectProperties) SetSummary(v vocab.PropertySummary) {
	o.PropertySummary = v
}

func (o *ObjectProperties) GetId() vocab.PropertyId {
	return o.PropertyId
}

func (o *ObjectProperties) SetId(v vocab.PropertyId) {
	o.PropertyId = v
}
