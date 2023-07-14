package impl

import (
	"fmt"
	"github.com/Grady-Saccullo/activity-pub-go/internal/types/json_ld"
	"github.com/Grady-Saccullo/activity-pub-go/internal/types/w3c/activity_stream/vocab"
)

const ObjectNoteTypeValue = "Note"

type ObjectNote struct {
	ObjectProperties
	alias *string
}

func DeserializeObjectNote(d map[string]interface{}, ldAliases map[string]string) (vocab.ObjectNote, error) {
	alias := json_ld.GetJsonLDContext(ldAliases, "https://www.w3.org/ns/activitystreams")

	s, ok := json_ld.GetType(d, alias)

	if !ok {
		return nil, fmt.Errorf("type is not defined")
	} else if *s != ObjectNoteTypeValue {
		return nil, nil
	}

	ret := ObjectNote{}

	if err := deserializeObjectProperties(d, ldAliases, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}
