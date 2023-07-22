package activity_streams_v2_impl

import (
	"github.com/Grady-Saccullo/activity-pub-go/pkg/w3c/activity_streams/v2/vocab"
	"github.com/Grady-Saccullo/activity-pub-go/pkg/w3c/json_ld/v1"
)

type ObjectType struct {
	alias          *string
	ObjectTypeNote activity_streams_v2_vocab.ObjectNote
}

func DeserializeObject(d map[string]interface{}, ldAliases map[string]string) (activity_streams_v2_vocab.Object, error) {
	alias := json_ld_v1.GetJsonLDContext(ldAliases, "https://www.w3.org/ns/activitystreams")

	t, ok := json_ld_v1.GetType(d, alias)
	if !ok {
		return nil, nil
	}

	switch *t {
	case ObjectNoteTypeValue:
		if v, err := DeserializeObjectNote(d, ldAliases); err != nil {
			return nil, err
		} else {
			return &ObjectType{
				alias:          alias,
				ObjectTypeNote: v,
			}, nil
		}
	}

	return nil, nil
}

func (o *ObjectType) GetObjectNote() activity_streams_v2_vocab.ObjectNote {
	return o.ObjectTypeNote
}
