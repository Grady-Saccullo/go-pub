package impl

import (
	json_ld_v1 "github.com/Grady-Saccullo/go-pub/pkg/w3c/json_ld/v1/helpers"
	"github.com/stretchr/testify/assert"
	"testing"
)

const testDeserializePropertyNameJson = `{
	"@context": "http://www.w3.org/ns/activitystreams",
	"type": "Note",
	"name": "Some Name"
}`

func TestDeserializePropertyName(t *testing.T) {
	data, ldAliases, err := json_ld_v1.ParseJson([]byte(testDeserializePropertyNameJson))
	if err != nil {
		t.Fatal(err)
	}

	pn, err := DeserializePropertyName(data, ldAliases)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "Some Name", *pn.GetString())
}

const testDeserializePropertyNameJson_MissingIdKey = `{
	"@context": "http://www.w3.org/ns/activitystreams",
	"type": "Note"
}`

func TestDeserializePropertyName_MissingIdKey(t *testing.T) {
	data, ldAliases, err := json_ld_v1.ParseJson([]byte(testDeserializePropertyNameJson_MissingIdKey))
	if err != nil {
		t.Fatal(err)
	}

	pn, err := DeserializePropertyName(data, ldAliases)
	if err != nil {
		t.Fatal(err)
	}
	assert.Nil(t, pn.GetString())
}

const testDeserializePropertyName_NameAndNameMapJson = `{
	"@context": "http://www.w3.org/ns/activitystreams",
	"type": "Note",
	"name": "Some Name",
	"nameMap": {
		"en": "Some Name"
	}
}`

func TestDeserializePropertyName_NameAndNameMap(t *testing.T) {
	data, ldAliases, err := json_ld_v1.ParseJson([]byte(testDeserializePropertyName_NameAndNameMapJson))
	if err != nil {
		t.Fatal(err)
	}

	pn, err := DeserializePropertyName(data, ldAliases)
	if err != nil && pn == nil {
		return
	}
	t.Fatal("failed to validate json correctly")
}
