package helpers

const KeyType = "type"

func GetType(d map[string]interface{}, alias *string) (*string, bool) {
	if v, ok := GetProperty(d, alias, KeyType); ok {
		if s, ok := v.(string); ok {
			return &s, true
		}
	}

	return nil, false
}
