package vocab

type ObjectPropertyGetters interface {
	GetName() PropertyName
	GetSummary() PropertySummary
	GetId() PropertyId
}

type ObjectPropertySetters interface {
	SetName(PropertyName)
	SetSummary(PropertySummary)
	SetId(PropertyId)
}
