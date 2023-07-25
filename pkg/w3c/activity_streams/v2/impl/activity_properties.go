package impl

import "github.com/Grady-Saccullo/go-pub/pkg/w3c/activity_streams/v2/vocab"

type ActivityProperties struct {
	ObjectProperties
	object []vocab.PropertyObject
	actor  []vocab.PropertyActor
	target []vocab.PropertyTarget
}

func deserializeActivityProperties(d map[string]interface{}, ldAliases map[string]string, i vocab.ActivityPropertySetters) error {

	if err := deserializeObjectProperties(d, ldAliases, i); err != nil {
		return err
	}

	if v, err := DeserializePropertyObject(d, ldAliases); err != nil {
		return err
	} else if v != nil {
		i.SetObject(v)
	}

	if v, err := DeserializePropertyActor(d, ldAliases); err != nil {
		return err
	} else if v != nil {
		i.SetActor(v)
	}

	if v, err := DeserializePropertyTarget(d, ldAliases); err != nil {
		return err
	} else if v != nil {
		i.SetTarget(v)
	}

	return nil
}

func (a *ActivityProperties) GetObject() []vocab.PropertyObject {
	return a.object
}

func (a *ActivityProperties) SetObject(v []vocab.PropertyObject) {
	a.object = v
}

func (a *ActivityProperties) GetActor() []vocab.PropertyActor {
	return a.actor
}

func (a *ActivityProperties) SetActor(v []vocab.PropertyActor) {
	a.actor = v
}

func (a *ActivityProperties) GetTarget() []vocab.PropertyTarget {
	return a.target
}

func (a *ActivityProperties) SetTarget(v []vocab.PropertyTarget) {
	a.target = v
}
