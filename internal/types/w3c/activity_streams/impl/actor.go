package impl

import (
	"github.com/Grady-Saccullo/activity-pub-go/internal/types/json_ld"
	"github.com/Grady-Saccullo/activity-pub-go/internal/types/w3c/activity_streams/vocab"
)

type Actor struct {
	alias            *string
	ActorPerson      vocab.ActorPerson
	ActorApplication vocab.ActorApplication
	ActorGroup       vocab.ActorGroup
}

func DeserializeActor(d map[string]interface{}, ldAliases map[string]string) (vocab.Actor, error) {
	alias := json_ld.GetJsonLDContext(ldAliases, "https://www.w3.org/ns/activitystreams")

	t, ok := json_ld.GetType(d, alias)
	if !ok {
		return nil, nil
	}

	switch *t {
	case ActorPersonTypeValue:
		if v, err := DeserializeActorPerson(d, ldAliases); err != nil {
			return nil, err
		} else if v != nil {
			return &Actor{
				alias:       alias,
				ActorPerson: v,
			}, err
		}
	case ActorApplicationTypeValue:
		if v, err := DeserializeActorApplication(d, ldAliases); err != nil {
			return nil, err
		} else if v != nil {
			return &Actor{
				alias:            alias,
				ActorApplication: v,
			}, err
		}
	case ActorGroupTypeValue:
		if v, err := DeserializeActorGroup(d, ldAliases); err != nil {
			return nil, err
		} else if v != nil {
			return &Actor{
				alias:      alias,
				ActorGroup: v,
			}, err
		}
	}

	return nil, nil
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
