package vocab

type ActivityTypeCreate interface {
	GetPropertyObject() []PropertyObject

	GetPropertyName() PropertyName

	GetPropertySummary() PropertySummary
}
