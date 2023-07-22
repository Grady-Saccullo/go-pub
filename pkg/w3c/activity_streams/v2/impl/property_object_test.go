package activity_streams_v2_impl

import (
	json_ld_v1 "github.com/Grady-Saccullo/activity-pub-go/pkg/w3c/json_ld/v1"
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
		assert.Equal(t, objectJsonName, *i.GetObjectNote().GetPropertyName().GetString())
		assert.Equal(t, objectJsonSummary, *i.GetObjectNote().GetPropertySummary().GetString())
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

		if i.IsIRI() {
			assert.Equal(t, "https://exampel.com/v1/note", i.GetIRI().String())
		}
		if i.IsObject() {
			if i.GetObjectNote() != nil {
				assert.Equal(t, objectJsonName, *i.GetObjectNote().GetPropertyName().GetString())
				assert.Equal(t, objectJsonSummary, *i.GetObjectNote().GetPropertySummary().GetString())
			}
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
		if i.IsActivity() {
			if i.GetActivityAccept() != nil {
				for _, a := range i.GetActivityAccept().GetPropertyObject() {
					assert.Equal(t, "Going-Away Party for Jim", *a.GetActivityFollow().GetPropertyName().GetString())
				}
			}
		}
	}
}
