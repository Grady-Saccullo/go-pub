package json_ld_v1

import (
	"encoding/json"
	"fmt"
	"strings"
)

func ParseJson(j []byte) (data map[string]interface{}, ldAliases map[string]string, err error) {
	if err := json.Unmarshal(j, &data); err != nil {
		return nil, nil, err
	}

	ldContext, ok := data["@context"]
	if !ok {
		return nil, nil, fmt.Errorf("missing context")
	}

	ldAliases = toAliasMap(ldContext)
	return data, ldAliases, nil
}

// "mostly" borrow from github.com/go-fed/activity with some added modifications
// and changes. this only supports simple json-ld aliasing... for now this is ok
func toAliasMap(i interface{}) (m map[string]string) {
	m = make(map[string]string)
	toHttpHttpsFn := func(s string) (ok bool, http, https string) {
		if strings.HasPrefix(s, "http://") {
			ok = true
			http = s
			https = "https" + strings.TrimPrefix(s, "http")
		} else if strings.HasPrefix(s, "https://") {
			ok = true
			https = s
			http = "http" + strings.TrimPrefix(s, "https")
		}
		return
	}
	switch v := i.(type) {
	case string:
		// Single entry, no alias.
		if ok, http, https := toHttpHttpsFn(v); ok {
			m[http] = ""
			m[https] = ""
		} else {
			m[v] = ""
		}
	case []interface{}:
		// Recursively apply.
		for _, elem := range v {
			r := toAliasMap(elem)
			for k, val := range r {
				m[k] = val
			}
		}
	case map[string]interface{}:
		// Map any aliases.
		for k, val := range v {
			// Only handle string aliases.
			switch a := val.(type) {
			case string:
				if ok, http, https := toHttpHttpsFn(a); ok {
					m[http] = k
					m[https] = k
				} else {
					m[a] = k
				}
			}
		}
	}
	return
}
