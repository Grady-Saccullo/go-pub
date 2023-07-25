package impl

import (
	"github.com/Grady-Saccullo/go-pub/pkg/w3c/json_ld/v1/helpers"
	"github.com/stretchr/testify/assert"
	"testing"
)

const testDeserializeActivityCreateJson = `{
	"@context": "http://www.w3.org/ns/activitystreams",
	"type": "Create",
	"object": {
		"type": "Note",
		"name": "I am 1 level deep",
		"summary": "Just a troll"
	}
}`

func TestDeserializeActivityCreate(t *testing.T) {
	data, ldAliases, err := helpers.ParseJson([]byte(testDeserializeActivityCreateJson))
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

	if len(pn.GetObject()) == 0 {
		t.Fatal("failed to parse object into ")
	}

	for _, o := range pn.GetObject() {
		if n := o.GetNote(); o != nil {
			if na := n.GetName(); na != nil {
				assert.Equal(t, "I am 1 level deep", *na.GetString())

				na.SetString("I've been replaced")
				assert.Equal(t, "I've been replaced", *na.GetString())
			} else {
				t.Fatal("failed to parse name")
			}

			if s := n.GetSummary(); s != nil {
				assert.Equal(t, "Just a troll", *s.GetString())

				s.SetString("I've been replaced")
				assert.Equal(t, "I've been replaced", *s.GetString())
			} else {
				t.Fatal("failed to parse summary")
			}
		} else {
			t.Fatal("")
		}
	}
}
