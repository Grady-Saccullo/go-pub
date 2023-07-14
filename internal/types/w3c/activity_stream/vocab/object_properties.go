package vocab

type ObjectPropertyGetters interface {
	GetPropertyName() PropertyName
	GetPropertySummary() PropertySummary
	GetPropertyId() PropertyId
}

type ObjectPropertySetters interface {
	SetPropertyName(PropertyName)
	SetPropertySummary(PropertySummary)
	SetPropertyId(PropertyId)
}
