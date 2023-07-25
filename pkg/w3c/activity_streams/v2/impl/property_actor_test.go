package impl

import (
	json_ld_v1 "github.com/Grady-Saccullo/go-pub/pkg/w3c/json_ld/v1/helpers"
	"github.com/stretchr/testify/assert"
	"testing"
)

const testDeserializePropertyActor_IRIJson = `{
	"@context": "http://www.w3.org/ns/activitystreams",
	"type": "Note",
	"actor": ["https://w3.org/note"]
}`

func TestDeserializePropertyActor(t *testing.T) {
	data, ldAliases, err := json_ld_v1.ParseJson([]byte(testDeserializePropertyActor_IRIJson))
	if err != nil {
		t.Fatal(err)
	}

	pn, err := DeserializePropertyActor(data, ldAliases)
	if err != nil {
		t.Fatal(err)
	}

	for _, i := range pn {
		expected := "https://w3.org/note"
		assert.Equal(t, expected, i.GetIRI().String())
	}
}
