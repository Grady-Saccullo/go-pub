package impl

import (
	"fmt"
	"github.com/Grady-Saccullo/activity-pub-go/internal/types/json_ld"
	"github.com/Grady-Saccullo/activity-pub-go/internal/types/w3c/activity_streams/vocab"
)

const ActivityLikeTypeValue = "Like"

type ActivityLike struct {
	ActivityProperties
	alias *string
}

func DeserializeActivityLike(d map[string]interface{}, ldAliases map[string]string) (vocab.ActivityLike, error) {
	alias := json_ld.GetJsonLDContext(ldAliases, "https://www.w3.org/ns/activitystreams")

	s, ok := json_ld.GetType(d, alias)

	if !ok {
		return nil, fmt.Errorf("type is not defined")
	} else if *s != ActivityLikeTypeValue {
		return nil, nil
	}

	ret := ActivityLike{}

	if err := deserializeActivityProperties(d, ldAliases, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}