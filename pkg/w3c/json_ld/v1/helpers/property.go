package helpers

import "fmt"

func GetProperty(d map[string]interface{}, alias *string, name string) (interface{}, bool) {
	propertyName := name
	if alias != nil && len(*alias) > 0 {
		propertyName = fmt.Sprintf("%s:%s", *alias, name)
	}

	if v, ok := d[propertyName]; ok {
		return v, true
	}

	return nil, false
}
