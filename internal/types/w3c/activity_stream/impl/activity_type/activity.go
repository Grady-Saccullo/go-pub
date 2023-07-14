package activity_type

import (
	"github.com/Grady-Saccullo/activity-pub-go/internal/types/json_ld"
	"github.com/Grady-Saccullo/activity-pub-go/internal/types/w3c/activity_stream/impl/activity_type_create"
	"github.com/Grady-Saccullo/activity-pub-go/internal/types/w3c/activity_stream/vocab"
)

type ActivityType struct {
	alias              *string
	ActivityTypeCreate vocab.ActivityTypeCreate
}

func DeserializeActivityType(d map[string]interface{}, ldAliases map[string]string) (vocab.ActivityType, error) {
	alias := json_ld.GetJsonLDContext(ldAliases, "https://www.w3.org/ns/activitystreams")

	t, ok := json_ld.GetType(d, alias)
	if !ok {
		return nil, nil
	}

	switch *t {
	case activity_type_create.TypeValue:
		if v, err := mgr.DeserializeActivityTypeCreate()(d, ldAliases); err != nil {
			return nil, err
		} else if v != nil {
			return &ActivityType{
				alias:              alias,
				ActivityTypeCreate: v,
			}, err
		}
	}

	return nil, nil
}

func (a *ActivityType) GetActivityTypeCreate() vocab.ActivityTypeCreate {
	return a.ActivityTypeCreate
}
