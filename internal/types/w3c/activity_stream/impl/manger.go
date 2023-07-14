package impl

import (
	"github.com/Grady-Saccullo/activity-pub-go/internal/types/w3c/activity_stream/impl/activity_type"
	"github.com/Grady-Saccullo/activity-pub-go/internal/types/w3c/activity_stream/impl/activity_type_create"
	"github.com/Grady-Saccullo/activity-pub-go/internal/types/w3c/activity_stream/impl/object_type"
	"github.com/Grady-Saccullo/activity-pub-go/internal/types/w3c/activity_stream/impl/object_type_note"
	"github.com/Grady-Saccullo/activity-pub-go/internal/types/w3c/activity_stream/impl/property_name"
	"github.com/Grady-Saccullo/activity-pub-go/internal/types/w3c/activity_stream/impl/property_object"
	"github.com/Grady-Saccullo/activity-pub-go/internal/types/w3c/activity_stream/impl/property_summary"
	"github.com/Grady-Saccullo/activity-pub-go/internal/types/w3c/activity_stream/vocab"
)

type Manager struct {
}

var mgr *Manager

func init() {
	mgr = &Manager{}
	activity_type.SetManager(mgr)
	activity_type_create.SetManager(mgr)
	object_type.SetManager(mgr)
	object_type_note.SetManager(mgr)
	property_object.SetManager(mgr)
}

type Deserializers interface {
	DeserializeActivityType() func(map[string]interface{}, map[string]string) (vocab.ActivityType, error)
	DeserializeActivityTypeCreate() func(map[string]interface{}, map[string]string) (vocab.ActivityTypeCreate, error)
	DeserializeObjectType() func(map[string]interface{}, map[string]string) (vocab.ObjectType, error)
	DeserializeObjectTypeNote() func(map[string]interface{}, map[string]string) (vocab.ObjectTypeNote, error)
	DeserializePropertyName() func(map[string]interface{}, map[string]string) (vocab.PropertyName, error)
	DeserializePropertyObject() func(map[string]interface{}, map[string]string) ([]vocab.PropertyObject, error)
	DeserializePropertySummary() func(map[string]interface{}, map[string]string) (vocab.PropertySummary, error)
}

func (m *Manager) DeserializeActivityType() func(map[string]interface{}, map[string]string) (vocab.ActivityType, error) {
	return activity_type.DeserializeActivityType
}

func (m *Manager) DeserializeActivityTypeCreate() func(map[string]interface{}, map[string]string) (vocab.ActivityTypeCreate, error) {
	return activity_type_create.DeserializeActivityTypeCreate
}

func (m *Manager) DeserializeObjectType() func(map[string]interface{}, map[string]string) (vocab.ObjectType, error) {
	return object_type.DeserializeObjectType
}

func (m *Manager) DeserializeObjectTypeNote() func(map[string]interface{}, map[string]string) (vocab.ObjectTypeNote, error) {
	return object_type_note.DeserializeObjectTypeNote
}

func (m *Manager) DeserializePropertyName() func(map[string]interface{}, map[string]string) (vocab.PropertyName, error) {
	return property_name.DeserializePropertyName
}

func (m *Manager) DeserializePropertyObject() func(map[string]interface{}, map[string]string) ([]vocab.PropertyObject, error) {
	return property_object.DeserializePropertyObject
}

func (m *Manager) DeserializePropertySummary() func(map[string]interface{}, map[string]string) (vocab.PropertySummary, error) {
	return property_summary.DeserializePropertySummary
}
