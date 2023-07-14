package vocab

type ActivityPropertyGetters interface {
	ObjectPropertyGetters
	GetPropertyObject() []PropertyObject
	GetPropertyActor() []PropertyActor
	GetPropertyTarget() []PropertyTarget
}

type ActivityPropertySetters interface {
	ObjectPropertySetters
	SetPropertyObject([]PropertyObject)
	SetPropertyActor([]PropertyActor)
	SetPropertyTarget([]PropertyTarget)
}
