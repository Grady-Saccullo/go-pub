package vocab

type Actor interface {
	GetActorPerson() ActorPerson
	GetActorApplication() ActorApplication
	GetActorGroup() ActorGroup
}
