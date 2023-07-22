package impl

import (
	"github.com/Grady-Saccullo/activity-pub-go/internal/types/json_ld"
	"github.com/Grady-Saccullo/activity-pub-go/internal/types/w3c/activity_streams/vocab"
)

type Activity struct {
	alias          *string
	ActivityCreate vocab.ActivityCreate
	ActivityAccept vocab.ActivityAccept
	ActivityFollow vocab.ActivityFollow
	ActivityLike   vocab.ActivityLike
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
			return &Activity{
				alias:          alias,
				ActivityCreate: v,
			}, err
		}
	case ActivityAcceptTypeValue:
		if v, err := DeserializeActivityAccept(d, ldAliases); err != nil {
			return nil, err
		} else if v != nil {
			return &Activity{
				alias:          alias,
				ActivityAccept: v,
			}, err
		}
	case ActivityFollowTypeValue:
		if v, err := DeserializeActivityFollow(d, ldAliases); err != nil {
			return nil, err
		} else if v != nil {
			return &Activity{
				alias:          alias,
				ActivityFollow: v,
			}, err
		}
	case ActivityLikeTypeValue:
		if v, err := DeserializeActivityLike(d, ldAliases); err != nil {
			return nil, err
		} else if v != nil {
			return &Activity{
				alias:          alias,
				ActivityFollow: v,
			}, err
		}
	}

	return nil, nil
}

func (a *Activity) GetActivityCreate() vocab.ActivityCreate {
	return a.ActivityCreate
}

func (a *Activity) GetActivityAccept() vocab.ActivityAccept {
	return a.ActivityAccept
}

func (a *Activity) GetActivityFollow() vocab.ActivityFollow {
	return a.ActivityFollow
}

func (a *Activity) GetActivityLike() vocab.ActivityLike {
	return a.ActivityLike
}
