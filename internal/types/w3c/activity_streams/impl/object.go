package impl

import (
	"github.com/Grady-Saccullo/activity-pub-go/internal/types/json_ld"
	"github.com/Grady-Saccullo/activity-pub-go/internal/types/w3c/activity_streams/vocab"
)

type ObjectType struct {
	alias          *string
	ObjectTypeNote vocab.ObjectNote
}

func DeserializeObject(d map[string]interface{}, ldAliases map[string]string) (vocab.Object, error) {
	alias := json_ld.GetJsonLDContext(ldAliases, "https://www.w3.org/ns/activitystreams")

	t, ok := json_ld.GetType(d, alias)
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

func (o *ObjectType) GetObjectNote() vocab.ObjectNote {
	return o.ObjectTypeNote
}
