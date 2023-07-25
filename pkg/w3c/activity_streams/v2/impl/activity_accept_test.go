package impl

import (
	"github.com/Grady-Saccullo/go-pub/pkg/w3c/json_ld/v1/helpers"
	"github.com/stretchr/testify/assert"
	"testing"
)

const testDeserializeActivityAcceptJson = `{
	"@context": "http://www.w3.org/ns/activitystreams",
	"type": "Accept",
	"actor": "http://example.org/v1/account/john",
	"object": {
		"type": "Follow",
		"id": "https://example.com/v1/person/follow"
	}
}`

func TestDeserializeActivityAccept(t *testing.T) {
	data, ldAliases, err := helpers.ParseJson([]byte(testDeserializeActivityAcceptJson))
	if err != nil {
		t.Fatal(err)
	}

	pn, err := DeserializeActivityCreate(data, ldAliases)
	if err != nil {
		t.Fatal(err)
	}
	if pn == nil {
		return
	}

	for _, a := range pn.GetActor() {
		if iri := a.GetIRI(); iri != nil {
			assert.Equal(t, "http://example.org/v1/account/john", iri.String())
		} else {
			t.Fatal("failed to parse actor into iri")
		}
	}

	for _, o := range pn.GetObject() {
		if f := o.GetFollow(); f != nil {
			if n := f.GetName(); n != nil {
				assert.Equal(t, "https://example.com/v1/person/follow", *n.GetString())
			} else {
				t.Fatal("failed to get name")
			}
		} else {
			t.Fatal("failed to get follow")
		}
	}
}
