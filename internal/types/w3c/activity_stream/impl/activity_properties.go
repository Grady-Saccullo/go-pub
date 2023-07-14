package impl

import "github.com/Grady-Saccullo/activity-pub-go/internal/types/w3c/activity_stream/vocab"

type ActivityProperties struct {
	ObjectProperties
	PropertyObject []vocab.PropertyObject
	PropertyActor  []vocab.PropertyActor
	PropertyTarget []vocab.PropertyTarget
}

func deserializeActivityProperties(d map[string]interface{}, ldAliases map[string]string, i vocab.ActivityPropertySetters) error {
	if err := deserializeObjectProperties(d, ldAliases, i); err != nil {
		return err
	}

	if v, err := DeserializePropertyObject(d, ldAliases); err != nil {
		return err
	} else if v != nil {
		i.SetPropertyObject(v)
	}

	if v, err := DeserializePropertyActor(d, ldAliases); err != nil {
		return err
	} else if v != nil {
		i.SetPropertyActor(v)
	}

	if v, err := DeserializePropertyTarget(d, ldAliases); err != nil {
		return err
	} else if v != nil {
		i.SetPropertyTarget(v)
	}

	return nil
}

func (a *ActivityProperties) GetPropertyObject() []vocab.PropertyObject {
	return a.PropertyObject
}

func (a *ActivityProperties) SetPropertyObject(v []vocab.PropertyObject) {
	a.PropertyObject = v
}

func (a *ActivityProperties) GetPropertyActor() []vocab.PropertyActor {
	return a.PropertyActor
}

func (a *ActivityProperties) SetPropertyActor(v []vocab.PropertyActor) {
	a.PropertyActor = v
}

func (a *ActivityProperties) GetPropertyTarget() []vocab.PropertyTarget {
	return a.PropertyTarget
}

func (a *ActivityProperties) SetPropertyTarget(v []vocab.PropertyTarget) {
	a.PropertyTarget = v
}
