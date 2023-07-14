package json_ld

import "net/url"

func GetIRI(prop interface{}) (*url.URL, bool) {
	if v, ok := prop.(string); ok {
		if iri, err := url.Parse(v); err != nil {
			return nil, false
		} else {
			return iri, true
		}
	}

	return nil, false
}
