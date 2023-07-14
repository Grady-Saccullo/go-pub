package property_name

import (
	"github.com/Grady-Saccullo/activity-pub-go/internal/types/json_ld"
	"github.com/Grady-Saccullo/activity-pub-go/internal/types/w3c/activity_stream/vocab"
	"github.com/stretchr/testify/assert"
	"testing"
)

const nameJson = `{
	"@context": "http://www.w3.org/ns/activitystreams",
	"type": "Note",
	"name": "Some Note Title"
}`

var nameJsonExpected = "Some Note Title"

func baseDeserializePropertyName(t *testing.T) vocab.PropertyName {
	data, ldAliases, err := json_ld.ParseJson([]byte(nameJson))
	if err != nil {
		t.Fatal(err)
	}
	v, err := DeserializePropertyName(data, ldAliases)
	if err != nil {
		t.Fatal(err)
	}

	return v
}

func TestDeserializePropertyName(t *testing.T) {
	baseDeserializePropertyName(t)
}

func TestPropertyName_IsString(t *testing.T) {
	pn := baseDeserializePropertyName(t)

	assert.Equal(t, true, pn.IsString())
	assert.Equal(t, false, pn.IsRDFLangStringMap())
}

func TestPropertyName_GetString(t *testing.T) {
	pn := baseDeserializePropertyName(t)
	assert.Equal(t, nameJsonExpected, *pn.GetString())
}

func TestPropertyName_SetString(t *testing.T) {
	nameNewValue := "I am a new name"
	pn := baseDeserializePropertyName(t)

	assert.Equal(t, nameJsonExpected, *pn.GetString())

	pn.SetString(nameNewValue)

	assert.Equal(t, nameNewValue, *pn.GetString())
}

const nameMapJson = `{
	"@context": "http://www.w3.org/ns/activitystreams",
	"type": "Note",
	"nameMap": {
		"en": "Some Note Title"
	}
}`

var nameMapJsonExpected = map[string]string{
	"en": "Some Note Title",
}

func baseDeserializePropertyNameMap() (vocab.PropertyName, error) {
	data, ldAliases, err := json_ld.ParseJson([]byte(nameMapJson))
	if err != nil {
		return nil, err
	}

	return DeserializePropertyName(data, ldAliases)
}

func TestDeserializePropertyName2(t *testing.T) {
	_, err := baseDeserializePropertyNameMap()
	if err != nil {
		t.Fatal(err)
	}
	return
}

func TestPropertyName_IsRDFLangStringMap(t *testing.T) {
	pn, err := baseDeserializePropertyNameMap()
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, true, pn.IsRDFLangStringMap())
	assert.Equal(t, false, pn.IsString())
}

func TestPropertyName_GetRDFLangStringMap(t *testing.T) {
	pn, err := baseDeserializePropertyNameMap()
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, nameMapJsonExpected, *pn.GetRDFLangStringMap())
}

func TestPropertyName_SetRDFLangStringMap(t *testing.T) {
	nameNewValue := map[string]string{
		"en": "I am new text",
	}
	pn := baseDeserializePropertyName(t)

	assert.Equal(t, nameJsonExpected, *pn.GetString())

	pn.SetRDFLangStringMap(nameNewValue)

	assert.Equal(t, nameNewValue, *pn.GetRDFLangStringMap())
}

const nameUnknownJson = `{
	"@context": "http://www.w3.org/ns/activitystreams",
	"type": "Note",
	"nameMap": {
		"en": [ "Some Note Title" ]
	}
}`

func baseDeserializePropertyUnknown() (vocab.PropertyName, error) {
	data, ldAliases, err := json_ld.ParseJson([]byte(nameUnknownJson))
	if err != nil {
		return nil, err
	}

	return DeserializePropertyName(data, ldAliases)
}

func TestPropertyName_GetUnknown(t *testing.T) {
	pn, err := baseDeserializePropertyUnknown()
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, true, pn.IsUnknown())
}

func TestPropertyName_IsUnknown(t *testing.T) {
	pn, err := baseDeserializePropertyUnknown()
	if err != nil {
		t.Fatal(err)
	}

	expected := map[string]interface{}{
		"en": []interface{}{"Some Note Title"},
	}

	assert.Equal(t, expected, pn.GetUnknown())
}
