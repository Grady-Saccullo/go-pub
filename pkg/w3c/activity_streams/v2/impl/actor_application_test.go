package impl

import (
	json_ld_v1 "github.com/Grady-Saccullo/go-pub/pkg/w3c/json_ld/v1/helpers"
	"github.com/stretchr/testify/assert"
	"testing"
)

const testDeserializeActorApplicationJson = `{
	"@context": "http://www.w3.org/ns/activitystreams",
	"type": "Application",
	"id": "https://example.com/v1"
}`

func TestDeserializeActorApplication(t *testing.T) {
	data, ldAliases, err := json_ld_v1.ParseJson([]byte(testDeserializeActorApplicationJson))
	if err != nil {
		t.Fatal(err)
	}

	pn, err := DeserializeActorApplication(data, ldAliases)
	if err != nil {
		t.Fatal(err)
	}

	if pn == nil {
		return
	}

	if id := pn.GetId(); id != nil {
		assert.Equal(t, "https://example.com/v1", id.GetIRI().String())
	}
}
