package vocab

type ObjectTypeNote interface {
	GetPropertyName() PropertyName
	GetPropertySummary() PropertySummary
	GetPropertyObject() []PropertyObject
}
