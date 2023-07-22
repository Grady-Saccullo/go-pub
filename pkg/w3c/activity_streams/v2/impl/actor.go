package activity_streams_v2_impl

import (
	"github.com/Grady-Saccullo/activity-pub-go/pkg/w3c/activity_streams/v2/vocab"
	"github.com/Grady-Saccullo/activity-pub-go/pkg/w3c/json_ld/v1"
)

type Actor struct {
	alias            *string
	ActorPerson      activity_streams_v2_vocab.ActorPerson
	ActorApplication activity_streams_v2_vocab.ActorApplication
	ActorGroup       activity_streams_v2_vocab.ActorGroup
}

func DeserializeActor(d map[string]interface{}, ldAliases map[string]string) (activity_streams_v2_vocab.Actor, error) {
	alias := json_ld_v1.GetJsonLDContext(ldAliases, "https://www.w3.org/ns/activitystreams")

	t, ok := json_ld_v1.GetType(d, alias)
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

func (a *Actor) GetActorPerson() activity_streams_v2_vocab.ActorPerson {
	return a.ActorPerson
}

func (a *Actor) GetActorApplication() activity_streams_v2_vocab.ActorApplication {
	return a.ActorApplication
}

func (a *Actor) GetActorGroup() activity_streams_v2_vocab.ActorGroup {
	return a.ActorGroup
}
