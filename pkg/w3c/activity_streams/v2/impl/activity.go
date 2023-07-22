package activity_streams_v2_impl

import (
	"github.com/Grady-Saccullo/activity-pub-go/pkg/w3c/activity_streams/v2/vocab"
	json_ld_v1 "github.com/Grady-Saccullo/activity-pub-go/pkg/w3c/json_ld/v1"
)

type Activity struct {
	alias          *string
	ActivityCreate activity_streams_v2_vocab.ActivityCreate
	ActivityAccept activity_streams_v2_vocab.ActivityAccept
	ActivityFollow activity_streams_v2_vocab.ActivityFollow
	ActivityLike   activity_streams_v2_vocab.ActivityLike
}

func DeserializeActivity(d map[string]interface{}, ldAliases map[string]string) (activity_streams_v2_vocab.Activity, error) {
	alias := json_ld_v1.GetJsonLDContext(ldAliases, "https://www.w3.org/ns/activitystreams")

	t, ok := json_ld_v1.GetType(d, alias)
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

func (a *Activity) GetActivityCreate() activity_streams_v2_vocab.ActivityCreate {
	return a.ActivityCreate
}

func (a *Activity) GetActivityAccept() activity_streams_v2_vocab.ActivityAccept {
	return a.ActivityAccept
}

func (a *Activity) GetActivityFollow() activity_streams_v2_vocab.ActivityFollow {
	return a.ActivityFollow
}

func (a *Activity) GetActivityLike() activity_streams_v2_vocab.ActivityLike {
	return a.ActivityLike
}
