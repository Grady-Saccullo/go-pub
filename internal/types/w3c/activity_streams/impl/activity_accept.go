package impl

import (
	"fmt"
	"github.com/Grady-Saccullo/activity-pub-go/internal/types/json_ld"
	"github.com/Grady-Saccullo/activity-pub-go/internal/types/w3c/activity_streams/vocab"
)

const ActivityAcceptTypeValue = "Accept"

type ActivityAccept struct {
	ActivityProperties
	alias *string
}

func DeserializeActivityAccept(d map[string]interface{}, ldAliases map[string]string) (vocab.ActivityAccept, error) {
	alias := json_ld.GetJsonLDContext(ldAliases, "https://www.w3.org/ns/activitystreams")

	s, ok := json_ld.GetType(d, alias)

	if !ok {
		return nil, fmt.Errorf("type is not defined")
	} else if *s != ActivityAcceptTypeValue {
		return nil, nil
	}

	ret := ActivityFollow{}

	if err := deserializeActivityProperties(d, ldAliases, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}
