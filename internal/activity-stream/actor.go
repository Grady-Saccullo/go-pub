package activityStream

type ActorType string

const (
	ActorTypeApplication  ActorType = "Application"
	ActorTypeGroup        ActorType = "Group"
	ActorTypeOrganization ActorType = "Organization"
	ActorTypePerson       ActorType = "Person"
	ActorTypeService      ActorType = "Service"
)

type Application struct {
	ObjectCore
	Type ActorType `json:"type"`
}

type Group struct {
	ObjectCore
	Type ActorType `json:"type"`
}

type Organization struct {
	ObjectCore
	Type ActorType `json:"type"`
}

type Person struct {
	ObjectCore
	Type ActorType `json:"type"`
}

type Service struct {
	ObjectCore
	Type ActorType `json:"type"`
}
