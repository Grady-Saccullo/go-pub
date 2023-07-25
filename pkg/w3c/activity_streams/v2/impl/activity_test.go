package impl

import (
	json_ld_v1 "github.com/Grady-Saccullo/go-pub/pkg/w3c/json_ld/v1/helpers"
	"github.com/stretchr/testify/assert"
	"testing"
)

const testDeserializeActivityJson = `{
	"@context": "http://www.w3.org/ns/activitystreams",
	"type": "Create",
	"object": {
		"type": "Note",
		"name": "git checkout -b me",
		"summary": "Just a troll note"
	}
}`

func TestDeserializeActivity(t *testing.T) {
	data, ldAliases, err := json_ld_v1.ParseJson([]byte(testDeserializeActivityJson))
	if err != nil {
		t.Fatal(err)
	}

	pn, err := DeserializeActivity(data, ldAliases)
	if err != nil {
		t.Fatal(err)
	}

	if pn == nil {
		return
	}

	for _, i := range pn.GetCreate().GetObject() {
		if n := i.GetNote(); n != nil {
			if n := n.GetName(); n != nil {
				assert.Equal(t, "git checkout -b me", *n.GetString())
			} else {
				t.Fatal("failed to get property name")
			}
		}
	}

	if pn.GetCreate() != nil {
		for _, i := range pn.GetCreate().GetObject() {
			if n := i.GetNote(); n != nil {
				if n.GetName() != nil {
					assert.Equal(t, "git checkout -b me", *n.GetName().GetString())
				} else {
					t.Fatal("missing property name")
				}

				n.GetName().SetString("I will def werk!")
				assert.Equal(t, "I will def werk!", *n.GetName().GetString())

				if n.GetSummary() != nil {
					if n.GetSummary().GetString() != nil {
						assert.Equal(t, "Just a troll note", *n.GetSummary().GetString())
					}
				} else {
					t.Fatal("missing property summary")
				}

				n.GetSummary().SetString("A new summary")
				assert.Equal(t, "A new summary", *n.GetSummary().GetString())
			}

		}
	}
}

const testActivity_GetLikeJson = `{
	"@context": "http://www.w3.org/ns/activitystreams",
	"type": "Like",
	"object": {
		"id": "https://example.com/v1/note",
		"type": "Note"
	}
}`

func TestActivity_GetLike(t *testing.T) {
	data, ldAliases, err := json_ld_v1.ParseJson([]byte(testActivity_GetLikeJson))
	if err != nil {
		t.Fatal(err)
	}

	pn, err := DeserializeActivity(data, ldAliases)
	if err != nil {
		t.Fatal(err)
	}

	if pn == nil {
		return
	}
	if l := pn.GetLike(); l != nil {
		for _, i := range l.GetObject() {
			if n := i.GetNote(); n != nil {
				assert.Equal(t, "https://example.com/v1/note", n.GetId().GetIRI().String())
			}
		}
	}
}
