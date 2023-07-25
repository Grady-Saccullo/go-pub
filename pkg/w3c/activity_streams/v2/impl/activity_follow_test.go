package impl

import (
	json_ld_v1 "github.com/Grady-Saccullo/go-pub/pkg/w3c/json_ld/v1/helpers"
	"github.com/stretchr/testify/assert"
	"testing"
)

const testDeserializeActivityFollowJson = `{
	"@context": "http://www.w3.org/ns/activitystreams",
	"type": "Follow",
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

func TestDeserializeActivityFollow(t *testing.T) {
	data, ldAliases, err := json_ld_v1.ParseJson([]byte(testDeserializeActivityFollowJson))
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

	for _, i := range pn.GetActor() {
		if n := i.GetNote(); n != nil {
			assert.Equal(t, "https://example.com/v1/account/john", n.GetId().GetIRI().String())
			assert.Equal(t, "John", *n.GetName().GetString())
		}
	}
}
