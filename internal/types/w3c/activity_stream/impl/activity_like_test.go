package impl

import (
	"github.com/Grady-Saccullo/activity-pub-go/internal/types/json_ld"
	"github.com/stretchr/testify/assert"
	"testing"
)

const testDeserializeActivityLikeJson = `{
	"@context": "http://www.w3.org/ns/activitystreams",
	"type": "Like",
	"object": {
		"type": "Note",
		"name": "John"
	},
	"actor": {
		"type": "Note",
		"id": "https://example.com/v1/account/john",
		"name": "John"
	}
}`

func TestDeserializeActivityLike(t *testing.T) {
	data, ldAliases, err := json_ld.ParseJson([]byte(testDeserializeActivityLikeJson))
	if err != nil {
		t.Fatal(err)
	}

	pn, err := DeserializeActivityFollow(data, ldAliases)
	if err != nil {
		t.Fatal(err)
	}

	if pn == nil {
		return
	}

	for _, i := range pn.GetPropertyActor() {
		if i.IsObject() {
			if n := i.GetObjectNote(); n != nil {
				assert.Equal(t, "https://example.com/v1/account/john", n.GetPropertyId().GetIRI().String())
				assert.Equal(t, "John", *n.GetPropertyName().GetString())
			}
		}
	}
}
