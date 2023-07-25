package impl

import (
	json_ld_v1 "github.com/Grady-Saccullo/go-pub/pkg/w3c/json_ld/v1/helpers"
	"github.com/stretchr/testify/assert"
	"testing"
)

const testDeserializePropertyObject_IRIJson = `{
	"@context": "http://www.w3.org/ns/activitystreams",
	"type": "Note",
	"object": "https://w3.org/note"
}`

func TestDeserializePropertyObject_IRI(t *testing.T) {
	data, ldAliases, err := json_ld_v1.ParseJson([]byte(testDeserializePropertyObject_IRIJson))
	if err != nil {
		t.Fatal(err)
	}

	pn, err := DeserializePropertyObject(data, ldAliases)
	if err != nil {
		t.Fatal(err)
	}

	for _, i := range pn {
		expected := "https://w3.org/note"
		assert.Equal(t, expected, i.GetIRI().String())
	}
}

const testDeserializePropertyObjectJsonSingle = `{
	"@context": "http://www.w3.org/ns/activitystreams",
	"type": "Create",
	"object": {
		"type": "Note",
		"name": "I prolly work...",
		"summary": "Just a troll"
	}
}`

const objectJsonName = "I prolly work..."
const objectJsonSummary = "Just a troll"

func TestDeserializePropertyObject_Single(t *testing.T) {
	data, ldAliases, err := json_ld_v1.ParseJson([]byte(testDeserializePropertyObjectJsonSingle))
	if err != nil {
		t.Fatal(err)
	}

	pn, err := DeserializePropertyObject(data, ldAliases)
	if err != nil {
		t.Fatal(err)
	}

	for _, i := range pn {
		assert.Equal(t, objectJsonName, *i.GetNote().GetName().GetString())
		assert.Equal(t, objectJsonSummary, *i.GetNote().GetSummary().GetString())
	}
}

const testDeserializePropertyObject_ArrayJson = `{
	"@context": "http://www.w3.org/ns/activitystreams",
	"type": "Create",
	"object": [
		"https://exampel.com/v1/note",
		{
			"type": "Note",
			"name": "I prolly work...",
			"summary": "Just a troll"
		},
		{
			"type": "Note",
			"name": "I prolly work...",
			"summary": "Just a troll"
		},
		{
			"type": "Note",
			"name": "I prolly work...",
			"summary": "Just a troll"
		},
		{
			"type": "Note",
			"name": "I prolly work...",
			"summary": "Just a troll"
		},
		{
			"type": "Note",
			"name": "I prolly work...",
			"summary": "Just a troll"
		},
		{
			"type": "Note",
			"name": "I prolly work...",
			"summary": "Just a troll"
		}
	]
}`

func TestDeserializePropertyObject_Array(t *testing.T) {
	data, ldAliases, err := json_ld_v1.ParseJson([]byte(testDeserializePropertyObject_ArrayJson))
	if err != nil {
		t.Fatal(err)
	}

	pn, err := DeserializePropertyObject(data, ldAliases)
	if err != nil {
		t.Fatal(err)
	}

	for _, i := range pn {

		if i.GetIRI() != nil {
			assert.Equal(t, "https://exampel.com/v1/note", i.GetIRI().String())
		}
		if n := i.GetNote(); n != nil {
			assert.Equal(t, objectJsonName, *n.GetName().GetString())
			assert.Equal(t, objectJsonSummary, *n.GetSummary().GetString())
		}
	}
}

const testDeserializePropertyObject_NestedJson = `{
	"@context": "http://www.w3.org/ns/activitystreams",
	"type": "Create",
  	"object": {
    	"type": "Accept",
    	"actor": "http://example.org/v1/account/john",
		"object": {
			"type": "Follow",
			"name": "Going-Away Party for Jim"
		}
  	}
}`

func TestDeserializePropertyObject_Nested(t *testing.T) {
	data, ldAliases, err := json_ld_v1.ParseJson([]byte(testDeserializePropertyObject_NestedJson))
	if err != nil {
		t.Fatal(err)
	}

	pn, err := DeserializePropertyObject(data, ldAliases)
	if err != nil {
		t.Fatal(err)
	}

	for _, i := range pn {
		if n := i.GetNote(); n != nil {
			assert.Equal(t, objectJsonName, *n.GetName().GetString())
			assert.Equal(t, objectJsonSummary, *n.GetSummary().GetString())
		}
		if a := i.GetAccept(); a != nil {
			for _, o := range a.GetObject() {
				assert.Equal(t, "Going-Away Party for Jim", *o.GetFollow().GetName().GetString())
			}
		}
	}
}
