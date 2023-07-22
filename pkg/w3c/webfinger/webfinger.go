package webfinger

type WebFinger struct {
	Subject    string                 `json:"subject"`
	Aliases    []string               `json:"aliases,omitempty"`
	Properties map[string]interface{} `json:"properties,omitempty"`
	Links      []Link                 `json:"links"`
	Expires    string                 `json:"expires,omitempty"`
}

type Link struct {
	Rel        string                 `json:"rel"`
	Type       string                 `json:"type,omitempty"`
	Href       string                 `json:"href,omitempty"`
	Properties map[string]interface{} `json:"properties,omitempty"`
	Titles     map[string]string      `json:"titles,omitempty"`
}
