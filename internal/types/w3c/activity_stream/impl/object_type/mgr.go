package object_type

import "github.com/Grady-Saccullo/activity-pub-go/internal/types/w3c/activity_stream/vocab"

var mgr localManger

type localManger interface {
	DeserializeObjectTypeNote() func(map[string]interface{}, map[string]string) (vocab.ObjectTypeNote, error)
}

func SetManager(m localManger) {
	mgr = m
}
