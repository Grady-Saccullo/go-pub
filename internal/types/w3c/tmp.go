package w3c

import (
	"github.com/Grady-Saccullo/activity-pub-go/internal/types/json_ld"
	"log"
)

type TypeNote struct {
}

func DeserializeTypeNote(d map[string]interface{}, ldAliases map[string]string) (*TypeNote, error) {
	alias := json_ld.GetJsonLDContext(ldAliases, "https://www.w3.org/ns/activitystreams")

	prop, ok := json_ld.GetProperty(d, alias, "Note")
	if !ok {
		return nil, nil
	}

	if list, ok := prop.([]interface{}); ok {
		// de-serialize each piece of note
		log.Println(list)
	} else {
		// de-serialize either the string...
		log.Println(prop)
	}

	return nil, nil
}

const propertyTypeAttachment = "Attachment"

func DeserializePropertyAttachment(d map[string]interface{}, ldAliases map[string]string) (interface{}, error) {
	alias := json_ld.GetJsonLDContext(ldAliases, "https://www.w3.org/ns/activitystreams")

	prop, ok := json_ld.GetProperty(d, alias, propertyTypeAttachment)
	if !ok {
		return nil, nil
	}

	// NOTE:
	// 	Items can be of type string (IRI), Activity/Object
	if list, ok := prop.([]interface{}); ok {
		// TODO: deserialize each item in the attachment by triaging to handler (need to figure out max depth)
		log.Println(list)
	} else {
		// TODO: deserialize the single item in the attachment by triaging to handler (need to figure out max depth)
		log.Println(prop)
	}

	return prop, nil
}
