package activity_streams_v2_impl

import (
	json_ld_v1 "github.com/Grady-Saccullo/activity-pub-go/pkg/w3c/json_ld/v1"
	"github.com/stretchr/testify/assert"
	"testing"
)

const testDeserializeActorPersonJson = `{
	"@context": "http://www.w3.org/ns/activitystreams",
	"type": "Group",
	"id": "https://example.com/v1"
}`

func TestDeserializeActorPerson(t *testing.T) {
	data, ldAliases, err := json_ld_v1.ParseJson([]byte(testDeserializeActorPersonJson))
	if err != nil {
		t.Fatal(err)
	}

	pn, err := DeserializeActorPerson(data, ldAliases)
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
