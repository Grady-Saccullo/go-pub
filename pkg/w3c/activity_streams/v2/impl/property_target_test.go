package impl

import (
	json_ld_v1 "github.com/Grady-Saccullo/go-pub/pkg/w3c/json_ld/v1/helpers"
	"github.com/stretchr/testify/assert"

	"testing"
)

const testDeserializePropertyTargetJson = `{
	"@context": "http://www.w3.org/ns/activitystreams",
	"type": "Note",
	"target": "https://w3.org/note"
}`

func TestDeserializePropertyTarget(t *testing.T) {
	data, ldAliases, err := json_ld_v1.ParseJson([]byte(testDeserializePropertyTargetJson))
	if err != nil {
		t.Fatal(err)
	}

	pn, err := DeserializePropertyTarget(data, ldAliases)
	if err != nil {
		t.Fatal(err)
	}
	for _, i := range pn {
		assert.Equal(t, "https://w3.org/note", i.GetIRI().String())
	}
}

const testDeserializePropertyTarget_ObjectJson = `{
	"@context": "http://www.w3.org/ns/activitystreams",
	"type": "Follow",
	"target": {
		"type": "Note",
		"id": "https://example.com/v1/person"
	}
}`

func TestDeserializePropertyTarget_Object(t *testing.T) {
	data, ldAliases, err := json_ld_v1.ParseJson([]byte(testDeserializePropertyTarget_ObjectJson))
	if err != nil {
		t.Fatal(err)
	}

	pn, err := DeserializePropertyTarget(data, ldAliases)
	if err != nil {
		t.Fatal(err)
	}

	for _, i := range pn {
		if n := i.GetNote(); n != nil {
			assert.Equal(t, "https://example.com/v1/person", n.GetId().GetIRI().String())
		} else {
			t.Fatal("failed to parse note")
		}
	}
}

const testDeserializePropertyTarget_ArrayJson = `{
	"@context": "http://www.w3.org/ns/activitystreams",
	"type": "Follow",
	"target": [{
		"type": "Note",
		"id": "https://example.com/v1/person"
	}]
}`

func TestDeserializePropertyTarget_Array(t *testing.T) {
	data, ldAliases, err := json_ld_v1.ParseJson([]byte(testDeserializePropertyTarget_ArrayJson))
	if err != nil {
		t.Fatal(err)
	}

	pn, err := DeserializePropertyTarget(data, ldAliases)
	if err != nil {
		t.Fatal(err)
	}

	for _, i := range pn {
		if n := i.GetNote(); n != nil {
			assert.Equal(t, "https://example.com/v1/person", n.GetId().GetIRI().String())
		} else {
			t.Fatal("failed to parse note")
		}
	}
}
