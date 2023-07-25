package vocab

type ActivityPropertyGetters interface {
	ObjectPropertyGetters
	GetObject() []PropertyObject
	GetActor() []PropertyActor
	GetTarget() []PropertyTarget
}

type ActivityPropertySetters interface {
	ObjectPropertySetters
	SetObject([]PropertyObject)
	SetActor([]PropertyActor)
	SetTarget([]PropertyTarget)
}
