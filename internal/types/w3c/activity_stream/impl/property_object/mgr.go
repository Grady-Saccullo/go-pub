package property_object

import "github.com/Grady-Saccullo/activity-pub-go/internal/types/w3c/activity_stream/vocab"

var mgr localManger

type localManger interface {
	DeserializeObjectType() func(map[string]interface{}, map[string]string) (vocab.ObjectType, error)
}

func SetManager(m localManger) {
	mgr = m
}
