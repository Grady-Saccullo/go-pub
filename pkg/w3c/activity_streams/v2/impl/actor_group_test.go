package impl

import (
	json_ld_v1 "github.com/Grady-Saccullo/go-pub/pkg/w3c/json_ld/v1/helpers"
	"github.com/stretchr/testify/assert"
	"testing"
)

const testDeserializeActorGroupJson = `{
	"@context": "http://www.w3.org/ns/activitystreams",
	"type": "Group",
	"id": "https://example.com/v1"
}`

func TestDeserializeActorGroup(t *testing.T) {
	data, ldAliases, err := json_ld_v1.ParseJson([]byte(testDeserializeActorGroupJson))
	if err != nil {
		t.Fatal(err)
	}

	pn, err := DeserializeActorGroup(data, ldAliases)
	if err != nil {
		t.Fatal(err)
	}

	if pn == nil {
		return
	}

	if pn.GetId() != nil {
		assert.Equal(t, "https://example.com/v1", pn.GetId().GetIRI().String())
	}
}
