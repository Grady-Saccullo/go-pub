package _old_

type LinkType string

const (
	LinkTypeLink    LinkType = "Link"
	LinkTypeMention LinkType = "Mention"
)

type Link struct {
	JsonLDContext JsonLDContext `json:"@context,omitempty"`
	Type          LinkType      `json:"type"`
	PropertyName
	PropertyNameMap
	PropertyHref
	PropertyHrefLang
	PropertyHeight
	PropertyMediaType
	PropertyPreview
	PropertyWidth
}

type Mention struct {
	Link
}
