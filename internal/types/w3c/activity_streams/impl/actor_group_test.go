package impl

import (
	"github.com/Grady-Saccullo/activity-pub-go/internal/types/json_ld"
	"github.com/stretchr/testify/assert"
	"testing"
)

const testDeserializeActorGroupJson = `{
	"@context": "http://www.w3.org/ns/activitystreams",
	"type": "Group",
	"id": "https://example.com/v1"
}`

func TestDeserializeActorGroup(t *testing.T) {
	data, ldAliases, err := json_ld.ParseJson([]byte(testDeserializeActorGroupJson))
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

	if pn.GetPropertyId() != nil {
		assert.Equal(t, "https://example.com/v1", pn.GetPropertyId().GetIRI().String())
	}
}
