package impl

import (
	"github.com/Grady-Saccullo/go-pub/pkg/w3c/activity_streams/v2/vocab"
	"github.com/Grady-Saccullo/go-pub/pkg/w3c/json_ld/v1/helpers"
)

type Activity struct {
	alias  *string
	create vocab.ActivityCreate
	accept vocab.ActivityAccept
	follow vocab.ActivityFollow
	like   vocab.ActivityLike
}

func DeserializeActivity(d map[string]interface{}, ldAliases map[string]string) (vocab.Activity, error) {
	alias := helpers.GetJsonLDContext(ldAliases, "https://www.w3.org/ns/activitystreams")

	t, ok := helpers.GetType(d, alias)
	if !ok {
		return nil, nil
	}

	ret := &Activity{
		alias:  alias,
		create: &ActivityCreate{},
		accept: &ActivityAccept{},
		follow: &ActivityFollow{},
		like:   &ActivityLike{},
	}

	switch *t {
	case ActivityCreateTypeValue:
		if v, err := DeserializeActivityCreate(d, ldAliases); err != nil {
			return nil, err
		} else if v != nil {
			ret.create = v
		}
	case ActivityAcceptTypeValue:
		if v, err := DeserializeActivityAccept(d, ldAliases); err != nil {
			return nil, err
		} else if v != nil {
			ret.accept = v
		}
	case ActivityFollowTypeValue:
		if v, err := DeserializeActivityFollow(d, ldAliases); err != nil {
			return nil, err
		} else if v != nil {
			ret.follow = v
		}
	case ActivityLikeTypeValue:
		if v, err := DeserializeActivityLike(d, ldAliases); err != nil {
			return nil, err
		} else if v != nil {
			ret.follow = v
		}
	}

	return ret, nil
}

func (a *Activity) GetCreate() vocab.ActivityCreate {
	return a.create
}

func (a *Activity) SetCreate(c vocab.ActivityCreate) {
	a.create = c
}

func (a *Activity) GetAccept() vocab.ActivityAccept {
	return a.accept
}

func (a *Activity) SetAccept(ac vocab.ActivityAccept) {
	a.accept = ac
}

func (a *Activity) GetFollow() vocab.ActivityFollow {
	return a.follow
}

func (a *Activity) SetFollow(f vocab.ActivityFollow) {
	a.follow = f
}

func (a *Activity) GetLike() vocab.ActivityLike {
	return a.like
}

func (a *Activity) SetLike(l vocab.ActivityLike) {
	a.like = l
}
