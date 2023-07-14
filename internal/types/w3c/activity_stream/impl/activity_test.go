package impl

import (
	"github.com/Grady-Saccullo/activity-pub-go/internal/types/json_ld"
	"github.com/stretchr/testify/assert"
	"testing"
)

const testDeserializeActivityJson = `{
	"@context": "http://www.w3.org/ns/activitystreams",
	"type": "Create",
	"object": {
		"type": "Note",
		"name": "git checkout -b me",
		"summary": "Just a troll note"
	}
}`

func TestDeserializeActivity(t *testing.T) {
	data, ldAliases, err := json_ld.ParseJson([]byte(testDeserializeActivityJson))
	if err != nil {
		t.Fatal(err)
	}

	pn, err := DeserializeActivity(data, ldAliases)

	if err != nil {
		t.Fatal(err)
	}

	if pn == nil {
		return
	}

	for _, i := range pn.GetActivityCreate().GetPropertyObject() {
		if i.GetObjectNote().GetPropertyName() != nil {
			assert.Equal(t, "git checkout -b me", *i.GetObjectNote().GetPropertyName().GetString())
		} else {
			t.Fatal("failed to get property name")
		}
	}

	if pn.GetActivityCreate() != nil {
		for _, i := range pn.GetActivityCreate().GetPropertyObject() {
			if i.GetObjectNote() != nil {
				if i.GetObjectNote().GetPropertyName() != nil {
					assert.Equal(t, "git checkout -b me", *i.GetObjectNote().GetPropertyName().GetString())
				} else {
					t.Fatal("missing property name")
				}

				i.GetObjectNote().GetPropertyName().SetString("I will def werk!")
				assert.Equal(t, "I will def werk!", *i.GetObjectNote().GetPropertyName().GetString())

				if i.GetObjectNote().GetPropertySummary() != nil {
					if i.GetObjectNote().GetPropertySummary().GetString() != nil {
						assert.Equal(t, "Just a troll note", *i.GetObjectNote().GetPropertySummary().GetString())
					}
				} else {
					t.Fatal("missing property summary")
				}

				i.GetObjectNote().GetPropertySummary().SetString("A new summary")
				assert.Equal(t, "A new summary", *i.GetObjectNote().GetPropertySummary().GetString())
			}
		}
	}
}
