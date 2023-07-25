package vocab

type PropertyName interface {
	GetString() *string

	SetString(string)

	GetRDFLangStringMap() *map[string]string

	SetRDFLangStringMap(map[string]string)
}
