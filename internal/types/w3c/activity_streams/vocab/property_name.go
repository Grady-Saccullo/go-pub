package vocab

type PropertyName interface {
	GetString() *string

	SetString(string)

	IsString() bool

	GetRDFLangStringMap() *map[string]string

	SetRDFLangStringMap(map[string]string)

	IsRDFLangStringMap() bool

	GetUnknown() interface{}

	IsUnknown() bool
}
