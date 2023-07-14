package activity_type

import (
	"github.com/Grady-Saccullo/activity-pub-go/internal/types/w3c/activity_stream/vocab"
)

var mgr localManger

type localManger interface {
	DeserializeActivityTypeCreate() func(map[string]interface{}, map[string]string) (vocab.ActivityTypeCreate, error)
}

func SetManager(m localManger) {
	mgr = m
}
