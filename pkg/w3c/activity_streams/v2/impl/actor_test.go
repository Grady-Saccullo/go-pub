package impl

import (
	json_ld_v1 "github.com/Grady-Saccullo/go-pub/pkg/w3c/json_ld/v1/helpers"
	"github.com/stretchr/testify/assert"
	"testing"
)

const testDeserializeActorJson = `{
	"@context": "http://www.w3.org/ns/activitystreams",
	"type": "Person",
	"id": "https://example.com/v1/person"
}`

const testDeserializeActor_WrongTypeJson = `{
	"@context": "http://www.w3.org/ns/activitystreams",
	"id": "https://example.com/v1/person"
}`

const testDeserializeActor_PersonJson = `{
	"@context": "http://www.w3.org/ns/activitystreams",
	"type": "Person",
	"id": "https://example.com/v1/person"
}`

const testDeserializeActor_ApplicationJson = `{
	"@context": "http://www.w3.org/ns/activitystreams",
	"type": "Application",
	"id": "https://example.com/v1/person"
}`

const testDeserializeActor_GroupJson = `{
	"@context": "http://www.w3.org/ns/activitystreams",
	"type": "Group",
	"id": "https://example.com/v1/person"
}`

func TestDeserializeActor(t *testing.T) {
	data, ldAliases, err := json_ld_v1.ParseJson([]byte(testDeserializeActorJson))
	if err != nil {
		t.Fatal(err)
	}

	pn, err := DeserializeActor(data, ldAliases)

	if pn == nil && err == nil {
		t.Fatal("failed to get type")
	}

	if pn == nil && err != nil {
		t.Fatal("failed to parse json")
	}

	if p := pn.GetActorPerson(); p != nil {
		assert.Equal(t, "https://example.com/v1/person", p.GetId().GetIRI().String())
	} else {
		t.Fatal("failed to parse actor")
	}
}

func TestDeserializeActor2(t *testing.T) {
	data, ldAliases, err := json_ld_v1.ParseJson([]byte(testDeserializeActor_WrongTypeJson))
	if err != nil {
		t.Fatal(err)
	}

	pn, err := DeserializeActor(data, ldAliases)

	if pn == nil && err == nil {
		return
	}

	t.Fatal("deserialized data with wrong type")
}

func TestActor_GetActorApplication(t *testing.T) {
	data, ldAliases, err := json_ld_v1.ParseJson([]byte(testDeserializeActor_ApplicationJson))
	if err != nil {
		t.Fatal(err)
	}

	pn, err := DeserializeActor(data, ldAliases)

	if err != nil {
		t.Fatal(err)
	}

	if pn == nil {
		t.Fatal("failed to parse json into actor")
	}

	if p := pn.GetActorApplication(); p != nil {
		assert.Equal(t, "https://example.com/v1/person", p.GetId().GetIRI().String())
	} else {
		t.Fatal("failed to parse actor application")
	}
}

func TestActor_GetActorPerson(t *testing.T) {
	data, ldAliases, err := json_ld_v1.ParseJson([]byte(testDeserializeActor_PersonJson))
	if err != nil {
		t.Fatal(err)
	}

	pn, err := DeserializeActor(data, ldAliases)
	if err != nil {
		t.Fatal(err)
	}

	if pn == nil {
		t.Fatal("failed to parse json into actor")
	}

	if p := pn.GetActorPerson(); p != nil {
		assert.Equal(t, "https://example.com/v1/person", p.GetId().GetIRI().String())
	} else {
		t.Fatal("failed to parse actor person")
	}
}

func TestActor_GetActorGroup(t *testing.T) {
	data, ldAliases, err := json_ld_v1.ParseJson([]byte(testDeserializeActor_GroupJson))
	if err != nil {
		t.Fatal(err)
	}

	pn, err := DeserializeActor(data, ldAliases)
	if err != nil {
		t.Fatal(err)
	}

	if pn == nil {
		t.Fatal("failed to parse json into actor")
	}

	if p := pn.GetActorGroup(); p != nil {
		assert.Equal(t, "https://example.com/v1/person", p.GetId().GetIRI().String())
	} else {
		t.Fatal("failed to parse actor group")
	}
}
