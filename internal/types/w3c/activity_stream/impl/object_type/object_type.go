package object_type

import (
	"github.com/Grady-Saccullo/activity-pub-go/internal/types/json_ld"
	"github.com/Grady-Saccullo/activity-pub-go/internal/types/w3c/activity_stream/impl/object_type_note"
	"github.com/Grady-Saccullo/activity-pub-go/internal/types/w3c/activity_stream/vocab"
)

type ObjectType struct {
	alias          *string
	ObjectTypeNote vocab.ObjectTypeNote
}

func DeserializeObjectType(d map[string]interface{}, ldAliases map[string]string) (vocab.ObjectType, error) {
	alias := json_ld.GetJsonLDContext(ldAliases, "https://www.w3.org/ns/activitystreams")

	t, ok := json_ld.GetType(d, alias)
	if !ok {
		return nil, nil
	}

	switch *t {
	case object_type_note.TypeValue:
		if v, err := mgr.DeserializeObjectTypeNote()(d, ldAliases); err != nil {
			return nil, err
		} else if v != nil {
			return &ObjectType{
				alias:          alias,
				ObjectTypeNote: v,
			}, err
		}
	}

	return nil, nil
}

func (o *ObjectType) GetObjectTypeNote() vocab.ObjectTypeNote {
	return o.ObjectTypeNote
}
