package impl

import (
	"github.com/Grady-Saccullo/activity-pub-go/internal/types/json_ld"
	"github.com/Grady-Saccullo/activity-pub-go/internal/types/w3c/activity_stream/vocab"
)

type ActivityType struct {
	alias              *string
	ActivityTypeCreate vocab.ActivityCreate
	ActivityAccept     vocab.ActivityAccept
	ActivityFollow     vocab.ActivityFollow
}

func DeserializeActivity(d map[string]interface{}, ldAliases map[string]string) (vocab.Activity, error) {
	alias := json_ld.GetJsonLDContext(ldAliases, "https://www.w3.org/ns/activitystreams")

	t, ok := json_ld.GetType(d, alias)
	if !ok {
		return nil, nil
	}

	switch *t {
	case ActivityCreateTypeValue:
		if v, err := DeserializeActivityCreate(d, ldAliases); err != nil {
			return nil, err
		} else if v != nil {
			return &ActivityType{
				alias:              alias,
				ActivityTypeCreate: v,
			}, err
		}
	case ActivityAcceptTypeValue:
		if v, err := DeserializeActivityAccept(d, ldAliases); err != nil {
			return nil, err
		} else if v != nil {
			return &ActivityType{
				alias:          alias,
				ActivityAccept: v,
			}, err
		}
	case ActivityFollowTypeValue:
		if v, err := DeserializeActivityFollow(d, ldAliases); err != nil {
			return nil, err
		} else if v != nil {
			return &ActivityType{
				alias:          alias,
				ActivityFollow: v,
			}, err
		}
	}

	return nil, nil
}

func (a *ActivityType) GetActivityCreate() vocab.ActivityCreate {
	return a.ActivityTypeCreate
}

func (a *ActivityType) GetActivityAccept() vocab.ActivityAccept {
	return a.ActivityAccept
}

func (a *ActivityType) GetActivityFollow() vocab.ActivityFollow {
	return a.ActivityFollow
}
