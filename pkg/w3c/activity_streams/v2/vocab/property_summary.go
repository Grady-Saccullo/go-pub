package activity_streams_v2_vocab

type PropertySummary interface {
	GetString() *string

	SetString(string)

	IsString() bool

	GetStringMap() *map[string]string

	SetStringMap(map[string]string)

	IsRDFLangStringMap() bool

	GetUnknown() interface{}

	IsUnknown() bool
}
