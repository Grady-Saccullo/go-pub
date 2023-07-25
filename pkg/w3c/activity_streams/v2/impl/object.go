package impl

import (
	"github.com/Grady-Saccullo/go-pub/pkg/w3c/activity_streams/v2/vocab"
	"github.com/Grady-Saccullo/go-pub/pkg/w3c/json_ld/v1/helpers"
)

type Object struct {
	alias *string
	note  vocab.ObjectNote
}

func DeserializeObject(d map[string]interface{}, ldAliases map[string]string) (vocab.Object, error) {
	alias := helpers.GetJsonLDContext(ldAliases, "https://www.w3.org/ns/activitystreams")

	t, ok := helpers.GetType(d, alias)
	if !ok {
		return nil, nil
	}

	ret := &Object{
		alias: alias,
		note:  nil,
	}
	switch *t {
	case ObjectNoteTypeValue:
		if v, err := DeserializeObjectNote(d, ldAliases); err != nil {
			return nil, err
		} else {
			ret.note = v
			return ret, nil
		}
	}

	return ret, nil
}

func (o *Object) GetNote() vocab.ObjectNote {
	return o.note
}

func (o *Object) SetNote(n vocab.ObjectNote) {
	o.note = n
}
