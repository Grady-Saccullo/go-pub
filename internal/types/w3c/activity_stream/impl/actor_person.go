package impl

import (
	"fmt"
	"github.com/Grady-Saccullo/activity-pub-go/internal/types/json_ld"
	"github.com/Grady-Saccullo/activity-pub-go/internal/types/w3c/activity_stream/vocab"
)

const ActorPersonTypeValue = "Person"

type ActorPerson struct {
	ActorProperties
	alias *string
}

func DeserializeActorPerson(d map[string]interface{}, ldAliases map[string]string) (vocab.ActorPerson, error) {
	alias := json_ld.GetJsonLDContext(ldAliases, "https://www.w3.org/ns/activitystreams")

	s, ok := json_ld.GetType(d, alias)

	if !ok {
		return nil, fmt.Errorf("type is not defined")
	} else if *s != ActorPersonTypeValue {
		return nil, nil
	}

	ret := ActorPerson{}

	if err := deserializeActorProperties(d, ldAliases, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}
