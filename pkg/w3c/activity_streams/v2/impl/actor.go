package impl

import (
	"github.com/Grady-Saccullo/go-pub/pkg/w3c/activity_streams/v2/vocab"
	"github.com/Grady-Saccullo/go-pub/pkg/w3c/json_ld/v1/helpers"
)

type Actor struct {
	alias            *string
	ActorPerson      vocab.ActorPerson
	ActorApplication vocab.ActorApplication
	ActorGroup       vocab.ActorGroup
}

func DeserializeActor(d map[string]interface{}, ldAliases map[string]string) (vocab.Actor, error) {
	alias := helpers.GetJsonLDContext(ldAliases, "https://www.w3.org/ns/activitystreams")

	t, ok := helpers.GetType(d, alias)
	if !ok {
		return nil, nil
	}

	ret := &Actor{
		alias: alias,
	}

	switch *t {
	case ActorPersonTypeValue:
		if v, err := DeserializeActorPerson(d, ldAliases); err != nil {
			return nil, err
		} else if v != nil {
			ret.ActorPerson = v
		}
	case ActorApplicationTypeValue:
		if v, err := DeserializeActorApplication(d, ldAliases); err != nil {
			return nil, err
		} else if v != nil {
			ret.ActorApplication = v
		}
	case ActorGroupTypeValue:
		if v, err := DeserializeActorGroup(d, ldAliases); err != nil {
			return nil, err
		} else if v != nil {
			ret.ActorGroup = v
		}
	}

	return ret, nil
}

func (a *Actor) GetActorPerson() vocab.ActorPerson {
	return a.ActorPerson
}

func (a *Actor) GetActorApplication() vocab.ActorApplication {
	return a.ActorApplication
}

func (a *Actor) GetActorGroup() vocab.ActorGroup {
	return a.ActorGroup
}
