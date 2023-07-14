package property_summary

import (
	"github.com/Grady-Saccullo/activity-pub-go/internal/types/json_ld"
	"github.com/stretchr/testify/assert"
	"testing"
)

const testDeserializePropertyNameJson1 = `{
	"@context": "http://www.w3.org/ns/activitystreams",
	"type": "Note",
	"summary": "I am a summary"
}`

func TestDeserializePropertySummary(t *testing.T) {
	data, ldAliases, err := json_ld.ParseJson([]byte(testDeserializePropertyNameJson1))
	if err != nil {
		t.Fatal(err)
	}

	pn, err := DeserializePropertySummary(data, ldAliases)
	if err != nil {
		t.Fatal(err)
	}

	expected := "I am a summary"

	assert.Equal(t, expected, *pn.GetString())
}

const testDeserializePropertyNameJson2 = `{
	"@context": "http://www.w3.org/ns/activitystreams",
	"type": "Note",
	"summaryMap": {
		"en": "I am a summary in english"
	}
}`

func TestDeserializePropertySummaryMap(t *testing.T) {
	data, ldAliases, err := json_ld.ParseJson([]byte(testDeserializePropertyNameJson2))
	if err != nil {
		t.Fatal(err)
	}

	pn, err := DeserializePropertySummary(data, ldAliases)
	if err != nil {
		t.Fatal(err)
	}

	expected := map[string]string{
		"en": "I am a summary in english",
	}

	assert.Equal(t, expected, *pn.GetStringMap())
}
