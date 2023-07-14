package impl

import (
	"fmt"
	"github.com/Grady-Saccullo/activity-pub-go/internal/types/json_ld"
	"github.com/Grady-Saccullo/activity-pub-go/internal/types/w3c/activity_stream/vocab"
)

const ActivityFollowTypeValue = "Follow"

type ActivityFollow struct {
	ActivityProperties
	alias *string
}

func DeserializeActivityFollow(d map[string]interface{}, ldAliases map[string]string) (vocab.ActivityFollow, error) {
	alias := json_ld.GetJsonLDContext(ldAliases, "https://www.w3.org/ns/activitystreams")

	s, ok := json_ld.GetType(d, alias)

	if !ok {
		return nil, fmt.Errorf("type is not defined")
	} else if *s != ActivityFollowTypeValue {
		return nil, nil
	}

	ret := ActivityFollow{}

	if err := deserializeActivityProperties(d, ldAliases, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}