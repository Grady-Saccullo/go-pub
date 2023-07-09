package activityStream

type LinkType string

const (
	LinkTypeLink    LinkType = "Link"
	LinkTypeMention LinkType = "Mention"
)

type Link struct {
	JsonLDContext JsonLDContext `json:"@context"`
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
