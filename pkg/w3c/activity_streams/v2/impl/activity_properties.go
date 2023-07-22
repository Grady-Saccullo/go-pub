package activity_streams_v2_impl

import (
	"github.com/Grady-Saccullo/activity-pub-go/pkg/w3c/activity_streams/v2/vocab"
)

type ActivityProperties struct {
	ObjectProperties
	PropertyObject []activity_streams_v2_vocab.PropertyObject
	PropertyActor  []activity_streams_v2_vocab.PropertyActor
	PropertyTarget []activity_streams_v2_vocab.PropertyTarget
}

func deserializeActivityProperties(d map[string]interface{}, ldAliases map[string]string, i activity_streams_v2_vocab.ActivityPropertySetters) error {
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

func (a *ActivityProperties) GetPropertyObject() []activity_streams_v2_vocab.PropertyObject {
	return a.PropertyObject
}

func (a *ActivityProperties) SetPropertyObject(v []activity_streams_v2_vocab.PropertyObject) {
	a.PropertyObject = v
}

func (a *ActivityProperties) GetPropertyActor() []activity_streams_v2_vocab.PropertyActor {
	return a.PropertyActor
}

func (a *ActivityProperties) SetPropertyActor(v []activity_streams_v2_vocab.PropertyActor) {
	a.PropertyActor = v
}

func (a *ActivityProperties) GetPropertyTarget() []activity_streams_v2_vocab.PropertyTarget {
	return a.PropertyTarget
}

func (a *ActivityProperties) SetPropertyTarget(v []activity_streams_v2_vocab.PropertyTarget) {
	a.PropertyTarget = v
}
