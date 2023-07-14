package impl

import (
	"github.com/Grady-Saccullo/activity-pub-go/internal/types/json_ld"
	"github.com/stretchr/testify/assert"
	"testing"
)

const TestManager_DeserializeObjectTypeNoteJson = `{
	"@context": "http://www.w3.org/ns/activitystreams",
	"type": "Note",
	"name": "I have a name!",
	"summary": "I am a summary"
}`

func TestManager_DeserializeObjectTypeNote(t *testing.T) {
	data, ldAliases, err := json_ld.ParseJson([]byte(TestManager_DeserializeObjectTypeNoteJson))
	if err != nil {
		t.Fatal(err)
	}

	pn, err := mgr.DeserializeObjectTypeNote()(data, ldAliases)
	if err != nil {
		t.Fatal(err)
	}
	if pn == nil {
		return
	}

	t.Log(pn)
	if pn.GetPropertySummary() != nil {
		if v := pn.GetPropertySummary().GetString(); v != nil {
		}
	}

	if pn.GetPropertyName() != nil {
		t.Log(*pn.GetPropertyName().GetString())
	}
}

const TestManager_DeserializeActivityTypeJson = `{
	"@context": "http://www.w3.org/ns/activitystreams",
	"type": "Create",
	"object": {
		"type": "Note",
		"name": "I am 1 level deep",
		"summary": "Just a troll",
		"object": {
			"type": "Note",
			"name": "I am 2 levels deep",
			"summary": "Just a troll",
			"object": {
				"type": "Note",
				"name": "I am 3 levels deep",
				"summary": "Just a troll"
			}
		}	
	}
}`

func TestManager_DeserializeActivityType(t *testing.T) {
	data, ldAliases, err := json_ld.ParseJson([]byte(TestManager_DeserializeActivityTypeJson))
	if err != nil {
		t.Fatal(err)
	}

	pn, err := mgr.DeserializeActivityType()(data, ldAliases)
	if err != nil {
		t.Fatal(err)
	}
	if pn == nil {
		return
	}

	for _, i := range pn.GetActivityTypeCreate().GetPropertyObject() {
		if i.GetObjectTypeNote().GetPropertyName() != nil {
			assert.Equal(t, "I am 1 level deep", *i.GetObjectTypeNote().GetPropertyName().GetString())
		}
		for _, i := range i.GetObjectTypeNote().GetPropertyObject() {
			if i.GetObjectTypeNote().GetPropertyName() != nil {
				assert.Equal(t, "I am 2 levels deep", *i.GetObjectTypeNote().GetPropertyName().GetString())
			}
			for _, i := range i.GetObjectTypeNote().GetPropertyObject() {
				if i.GetObjectTypeNote().GetPropertyName() != nil {
					t.Logf("lowest level %s", *i.GetObjectTypeNote().GetPropertyName().GetString())
					assert.Equal(t, "I am 3 levels deep", *i.GetObjectTypeNote().GetPropertyName().GetString())
				}
			}
		}
	}

	if pn.GetActivityTypeCreate() != nil {
		for _, i := range pn.GetActivityTypeCreate().GetPropertyObject() {
			if i.GetObjectTypeNote() != nil {
				if i.GetObjectTypeNote().GetPropertyName() != nil {
					assert.Equal(t, "I am 1 level deep", *i.GetObjectTypeNote().GetPropertyName().GetString())
				}

				i.GetObjectTypeNote().GetPropertyName().SetString("I will def werk!")
				assert.Equal(t, "I will def werk!", *i.GetObjectTypeNote().GetPropertyName().GetString())

				if i.GetObjectTypeNote().GetPropertySummary() != nil {
					if i.GetObjectTypeNote().GetPropertySummary().GetString() != nil {
						assert.Equal(t, "Just a troll", *i.GetObjectTypeNote().GetPropertySummary().GetString())
					}

					i.GetObjectTypeNote().GetPropertySummary().SetString("A new summary")
					assert.Equal(t, "A new summary", *i.GetObjectTypeNote().GetPropertySummary().GetString())
				}
			}
		}
	}
}

const testDeserializePropertyObjectJsonIRI = `{
	"@context": "http://www.w3.org/ns/activitystreams",
	"type": "Note",
	"object": "https://w3.org/note"
}`

func TestManager_DeserializePropertyObject_IRI(t *testing.T) {
	data, ldAliases, err := json_ld.ParseJson([]byte(testDeserializePropertyObjectJsonIRI))
	if err != nil {
		t.Fatal(err)
	}

	pn, err := mgr.DeserializePropertyObject()(data, ldAliases)
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

func TestManager_DeserializePropertyObject_Single(t *testing.T) {
	data, ldAliases, err := json_ld.ParseJson([]byte(testDeserializePropertyObjectJsonSingle))
	if err != nil {
		t.Fatal(err)
	}

	pn, err := mgr.DeserializePropertyObject()(data, ldAliases)
	if err != nil {
		t.Fatal(err)
	}

	for _, i := range pn {
		assert.Equal(t, objectJsonName, *i.GetObjectTypeNote().GetPropertyName().GetString())
		assert.Equal(t, objectJsonSummary, *i.GetObjectTypeNote().GetPropertySummary().GetString())
	}
}

const testDeserializePropertyObjectJsonArray = `{
	"@context": "http://www.w3.org/ns/activitystreams",
	"type": "Create",
	"object": [
		{
			"type": "Note",
			"name": "I prolly work...",
			"summary": "Just a troll"
		}
	]
}`

func TestManager_DeserializePropertyObject_Array(t *testing.T) {
	data, ldAliases, err := json_ld.ParseJson([]byte(testDeserializePropertyObjectJsonArray))
	if err != nil {
		t.Fatal(err)
	}

	pn, err := mgr.DeserializePropertyObject()(data, ldAliases)
	if err != nil {
		t.Fatal(err)
	}

	for _, i := range pn {
		assert.Equal(t, objectJsonName, *i.GetObjectTypeNote().GetPropertyName().GetString())
		assert.Equal(t, objectJsonSummary, *i.GetObjectTypeNote().GetPropertySummary().GetString())
	}
}
