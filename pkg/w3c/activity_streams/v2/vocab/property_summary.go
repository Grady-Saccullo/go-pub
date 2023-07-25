package vocab

type PropertySummary interface {
	GetString() *string

	SetString(string)

	GetStringMap() *map[string]string

	SetStringMap(map[string]string)

	GetUnknown() interface{}
}
