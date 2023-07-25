package impl

import (
	"github.com/Grady-Saccullo/go-pub/pkg/w3c/activity_streams/v2/vocab"
	json_ld_v1 "github.com/Grady-Saccullo/go-pub/pkg/w3c/json_ld/v1/helpers"
	"github.com/stretchr/testify/assert"
	"testing"
)

const testDeserializeObjectNoteJson = `{
	"@context": "http://www.w3.org/ns/activitystreams",
	"type": "Note",
	"name": "I have a name!",
	"summary": "I am a summary"
}`

func baseTestDeserializeObjectNote(t *testing.T) vocab.ObjectNote {
	data, ldAliases, err := json_ld_v1.ParseJson([]byte(testDeserializeObjectNoteJson))
	if err != nil {
		t.Fatal(err)
	}

	pn, err := DeserializeObjectNote(data, ldAliases)
	if err != nil {
		t.Fatal(err)
	}

	if pn == nil {
		t.Fatal("failed to parse object note")
	}
	return pn
}

func TestDeserializeObjectNote(t *testing.T) {
	baseTestDeserializeObjectNote(t)
}

func TestObjectNote_GetPropertyName(t *testing.T) {
	pn := baseTestDeserializeObjectNote(t)

	if pn.GetName() != nil {
		assert.Equal(t, "I have a name!", *pn.GetName().GetString())
	} else {
		t.Fatal("nil property name")
	}
}

func TestObjectNote_GetPropertySummary(t *testing.T) {
	pn := baseTestDeserializeObjectNote(t)

	if pn.GetSummary() != nil {
		assert.Equal(t, "I have a name!", *pn.GetName().GetString())
	} else {
		t.Fatal("nil property summary")
	}
}
