package rdf_v1

import (
	"fmt"
)

type LangString struct {
	LangString interface{} `json:"langString"`
}

func DeserializeLangString(d interface{}) (map[string]string, error) {
	o := make(map[string]string)
	if m, ok := d.(map[string]interface{}); ok {
		for k, v := range m {
			if s, ok := v.(string); ok {
				o[k] = s
			} else {
				return nil, fmt.Errorf("failed to parse rdf lang string at %s with value %v", k, v)
			}
		}
		return o, nil
	}
	return nil, fmt.Errorf("failed to parse lang string")
}

func SerializeLangString(d map[string]string) interface{} {
	return d
}
