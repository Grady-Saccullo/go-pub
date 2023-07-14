package _old_

type ActorType string

const (
	ActorTypeApplication  ActorType = "Application"
	ActorTypeGroup        ActorType = "Group"
	ActorTypeOrganization ActorType = "Organization"
	ActorTypePerson       ActorType = "Person"
	ActorTypeService      ActorType = "Service"
)

type Actor struct {
	ObjectCore
	Type ActorType `json:"type"`
}

type Application struct {
	Actor
}

type Group struct {
	Actor
}

type Organization struct {
	Actor
}

type Person struct {
	Actor
}

type Service struct {
	Actor
}
