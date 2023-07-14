package object_type_note

import "github.com/Grady-Saccullo/activity-pub-go/internal/types/w3c/activity_stream/vocab"

var mgr localManger

type localManger interface {
	DeserializePropertyName() func(map[string]interface{}, map[string]string) (vocab.PropertyName, error)
	DeserializePropertySummary() func(map[string]interface{}, map[string]string) (vocab.PropertySummary, error)
	DeserializePropertyObject() func(map[string]interface{}, map[string]string) ([]vocab.PropertyObject, error)
}

func SetManager(m localManger) {
	mgr = m
}
