package impl

import (
	json_ld_v1 "github.com/Grady-Saccullo/go-pub/pkg/w3c/json_ld/v1/helpers"
	"github.com/stretchr/testify/assert"
	"net/url"
	"testing"
)

const testDeserializePropertyIDJson = `{
	"@context": "http://www.w3.org/ns/activitystreams",
	"type": "Note",
	"id": "https://w3.org/note"
}`

func TestDeserializePropertyID(t *testing.T) {
	data, ldAliases, err := json_ld_v1.ParseJson([]byte(testDeserializePropertyIDJson))
	if err != nil {
		t.Fatal(err)
	}

	pn, err := DeserializePropertyId(data, ldAliases)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "https://w3.org/note", pn.GetIRI().String())
}

const testDeserializePropertyIDJson_MissingIdKey = `{
	"@context": "http://www.w3.org/ns/activitystreams",
	"type": "Note"
}`

func TestDeserializePropertyID_MissingIdKey(t *testing.T) {
	data, ldAliases, err := json_ld_v1.ParseJson([]byte(testDeserializePropertyIDJson_MissingIdKey))
	if err != nil {
		t.Fatal(err)
	}

	pn, err := DeserializePropertyId(data, ldAliases)
	if err != nil {
		t.Fatal(err)
	}

	assert.Nil(t, pn.GetIRI())
}

func TestDeserializePropertyId(t *testing.T) {
	data, ldAliases, err := json_ld_v1.ParseJson([]byte(testDeserializePropertyIDJson_MissingIdKey))
	if err != nil {
		t.Fatal(err)
	}

	pn, err := DeserializePropertyId(data, ldAliases)
	raw := "https://w3.org/note"
	u, err := url.Parse(raw)
	if err != nil {
		t.Fatal(err)
	}

	pn.SetIRI(u)
	assert.Equal(t, raw, pn.GetIRI().String())
}
