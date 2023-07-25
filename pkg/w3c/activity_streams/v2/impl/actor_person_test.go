package impl

import (
	json_ld_v1 "github.com/Grady-Saccullo/go-pub/pkg/w3c/json_ld/v1/helpers"
	"github.com/stretchr/testify/assert"
	"testing"
)

const testDeserializeActorPersonJson = `{
	"@context": "http://www.w3.org/ns/activitystreams",
	"type": "Person",
	"id": "https://example.com/v1"
}`

const testDeserializeActorPerson_WrongTypeJson = `{
	"@context": "http://www.w3.org/ns/activitystreams",
	"type": "Note",
	"id": "https://example.com/v1"
}`

const testDeserializeActorPerson_NoTypeJson = `{
	"@context": "http://www.w3.org/ns/activitystreams",
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

	if pn.GetId() != nil {
		assert.Equal(t, "https://example.com/v1", pn.GetId().GetIRI().String())
	}
}

func TestDeserializeActorPerson2(t *testing.T) {
	data, ldAliases, err := json_ld_v1.ParseJson([]byte(testDeserializeActorPerson_WrongTypeJson))
	if err != nil {
		t.Fatal(err)
	}

	pn, err := DeserializeActorPerson(data, ldAliases)

	if pn == nil && err == nil {
		return
	}

	t.Fatal("parsed invalid person")
}

func TestDeserializeActorPerson3(t *testing.T) {
	data, ldAliases, err := json_ld_v1.ParseJson([]byte(testDeserializeActorPerson_NoTypeJson))
	if err != nil {
		t.Fatal(err)
	}

	pn, err := DeserializeActorPerson(data, ldAliases)
	if pn == nil && err != nil {
		return
	}

	t.Fatal("parsed invalid person")
}
