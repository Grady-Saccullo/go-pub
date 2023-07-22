package activity_streams_v2_vocab

type Actor interface {
	GetActorPerson() ActorPerson
	GetActorApplication() ActorApplication
	GetActorGroup() ActorGroup
}
