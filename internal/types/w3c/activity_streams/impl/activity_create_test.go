package impl

import (
	"github.com/Grady-Saccullo/activity-pub-go/internal/types/json_ld"
	"github.com/stretchr/testify/assert"
	"testing"
)

const testDeserializeActivityCreateJson = `{
	"@context": "http://www.w3.org/ns/activitystreams",
	"type": "Create",
	"object": {
		"type": "Note",
		"name": "I am 1 level deep",
		"summary": "Just a troll"
	}
}`

func TestDeserializeActivityCreate(t *testing.T) {
	data, ldAliases, err := json_ld.ParseJson([]byte(testDeserializeActivityCreateJson))
	if err != nil {
		t.Fatal(err)
	}

	pn, err := DeserializeActivityCreate(data, ldAliases)
	if err != nil {
		t.Fatal(err)
	}
	if pn == nil {
		return
	}

	for _, i := range pn.GetPropertyObject() {
		if i.GetObjectNote() != nil {
			if i.GetObjectNote().GetPropertyName() != nil {
				assert.Equal(t, "I am 1 level deep", *i.GetObjectNote().GetPropertyName().GetString())
			}

			i.GetObjectNote().GetPropertyName().SetString("I will def werk!")
			assert.Equal(t, "I will def werk!", *i.GetObjectNote().GetPropertyName().GetString())

			if i.GetObjectNote().GetPropertySummary() != nil {
				if i.GetObjectNote().GetPropertySummary().GetString() != nil {
					assert.Equal(t, "Just a troll", *i.GetObjectNote().GetPropertySummary().GetString())
				}
			}

			i.GetObjectNote().GetPropertySummary().SetString("A new summary")
			assert.Equal(t, "A new summary", *i.GetObjectNote().GetPropertySummary().GetString())
		}
	}
}
