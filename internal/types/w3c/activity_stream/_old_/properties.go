package _old_

// JsonLDContext can of type string, array or map
type JsonLDContext interface{}

type PropertyId string

// PropertyActor inherits Object or Link
type PropertyActor struct {
	Actor interface{} `json:"actor,omitempty"`
}

// PropertyAttachment inherits Object or Link
type PropertyAttachment struct {
	Attachment interface{} `json:"attachment,omitempty"`
}

// PropertyAttributedTo inherits Actor, Object or Link
type PropertyAttributedTo struct {
	AttributedTo interface{} `json:"attributedTo,omitempty"`
}

// PropertyAudience inherits Object or Link
type PropertyAudience struct {
	Audience interface{} `json:"audience,omitempty"`
}

// PropertyBcc inherits Object or Link
type PropertyBcc struct {
	Bcc interface{} `json:"bcc,omitempty"`
}

// PropertyBto inherits Object or Link
type PropertyBto struct {
	Bto interface{} `json:"bto,omitempty"`
}

// PropertyCc inherits Object or Link
type PropertyCc struct {
	Cc interface{} `json:"cc,omitempty"`
}

// PropertyContext inherits Object or Link
type PropertyContext struct {
	Context interface{} `json:"context,omitempty"`
}

// PropertyCurrent inherits CollectionPage or Link
type PropertyCurrent struct {
	Current interface{} `json:"current,omitempty"`
}

// PropertyFirst inherits CollectionPage or Link as a functional type (single)
type PropertyFirst struct {
	Context interface{} `json:"context,omitempty"`
}

// PropertyGenerator inherits Object or Link
type PropertyGenerator struct {
	Generator interface{} `json:"generator,omitempty"`
}

// PropertyIcon inherits Image or Link
type PropertyIcon struct {
	Icon interface{} `json:"icon,omitempty"`
}

// PropertyImage inherits Image or Link
type PropertyImage struct {
	Image interface{} `json:"image,omitempty"`
}

// PropertyInReplyTo inherits Object or Link
type PropertyInReplyTo struct {
	InReplyTo interface{} `json:"inReplyTo,omitempty"`
}

// PropertyInstrument inherits Object or Link
type PropertyInstrument struct {
	Instrument interface{} `json:"instrument,omitempty"`
}

// PropertyLast inherits CollectionPage or Link  as a functional type (single)
type PropertyLast struct {
	Last interface{} `json:"last,omitempty"`
}

// PropertyLocation inherits Object or Link
type PropertyLocation struct {
	Location interface{} `json:"location,omitempty"`
}

// PropertyItems inherits Object or Link
type PropertyItems struct {
	Items interface{} `json:"items,omitempty"`
}

// PropertyOneOf inherits Object or Link
type PropertyOneOf struct {
	OneOf interface{} `json:"oneOf,omitempty"`
}

// PropertyClosed inherits Object, Link, date, or bool
type PropertyClosed struct {
	Closed interface{} `json:"closed,omitempty"`
}

// PropertyAnyOf inherits Object, Link, date, or bool
type PropertyAnyOf struct {
	AnyOf interface{} `json:"anyOf,omitempty"`
}

// PropertyOrigin inherits Object or Link
type PropertyOrigin struct {
	Origin interface{} `json:"origin,omitempty"`
}

// PropertyNext inherits CollectionPage or Link as a functional type (single)
type PropertyNext struct {
	Next interface{} `json:"next,omitempty"`
}

// PropertyObject inherits Object or Link
type PropertyObject struct {
	Object interface{} `json:"object,omitempty"`
}

// PropertyPrev inherits CollectionPage or Link as a functional type (single)
type PropertyPrev struct {
	Prev interface{} `json:"prev,omitempty"`
}

// PropertyPreview inherits Object or Link
type PropertyPreview struct {
	Preview interface{} `json:"preview,omitempty"`
}

// PropertyResult inherits Object or Link
type PropertyResult struct {
	Result interface{} `json:"result,omitempty"`
}

// PropertyReplies inherits Collection
type PropertyReplies struct {
	Replies interface{} `json:"replies,omitempty"`
}

// PropertyTarget inherits Object or Link
type PropertyTarget struct {
	Target interface{} `json:"target,omitempty"`
}

// PropertyTag inherits Object or Link
type PropertyTag struct {
	Tag interface{} `json:"tag,omitempty"`
}

// PropertyOrderedItems inherits Object or Link
type PropertyOrderedItems struct {
	OrderItems interface{} `json:"orderedItems"`
}

// PropertyPartOf inherits Collection or Link as a functional type (single)
type PropertyPartOf struct {
	PartOf interface{} `json:"partOf,omitempty"`
}

// PropertyRelationship inherits Object
type PropertyRelationship struct {
	Relationship interface{} `json:"relationship,omitempty"`
}

// PropertySubject inherits Object or Link as a functional type (single)
type PropertySubject struct {
	Subject interface{} `json:"subject,omitempty"`
}

// PropertyAccuracy is a functional type (single)
type PropertyAccuracy struct {
	Accuracy float32 `json:"accuracy,omitempty"`
}

// PropertyAltitude is a functional type (single)
type PropertyAltitude struct {
	Altitude float32 `json:"altitude,omitempty"`
}

// PropertyLatitude is a functional type (single)
type PropertyLatitude struct {
	Latitude float32 `json:"latitude,omitempty"`
}

// PropertyLongitude is a functional type (single)
type PropertyLongitude struct {
	Longitude float32 `json:"latitude,omitempty"`
}

// PropertyRadius is a functional type (single)
type PropertyRadius struct {
	Radius float32 `json:"radius,omitempty"`
}

type PropertyUnitsType string

const (
	PropertyUnitsTypeCm     PropertyUnitsType = "cm"
	PropertyUnitsTypeFeet   PropertyUnitsType = "feet"
	PropertyUnitsTypeInches PropertyUnitsType = "inches"
	PropertyUnitsTypeKm     PropertyUnitsType = "km"
	PropertyUnitsTypeM      PropertyUnitsType = "m"
	PropertyUnitsTypeMiles  PropertyUnitsType = "miles"
)

// PropertyUnits is a functional type (single)
type PropertyUnits struct {
	Units PropertyUnitsType `json:"units,omitempty"`
}

// PropertyDescribes inherits Object as a functional type (single)
type PropertyDescribes struct {
	Describes PropertyUnitsType `json:"describes,omitempty"`
}

// PropertyDeleted is a functional type (single)
type PropertyDeleted struct {
	Deleted string `json:"deleted,omitempty"`
}

// PropertyFormerType inherits Object
type PropertyFormerType struct {
	FormerType interface{} `json:"formerType,omitempty"`
}

// PropertyContent inherits Object
type PropertyContent struct {
	Content string `json:"content,omitempty"`
}

// PropertyContentMap map of lang to string
type PropertyContentMap struct {
	ContentMap map[string]string `json:"contentMap,omitempty"`
}

// PropertyName string
type PropertyName struct {
	Name string `json:"name,omitempty"`
}

// PropertyNameMap map of lang to string
type PropertyNameMap struct {
	NameMap map[string]string `json:"nameMap,omitempty"`
}

// PropertyEndTime string as a functional type (single)
type PropertyEndTime struct {
	EndTime string `json:"endTime,omitempty"`
}

// PropertyStartTime string as a functional type (single)
type PropertyStartTime struct {
	StartTime string `json:"startTime,omitempty"`
}

// PropertyPublished string as a functional type (single)
type PropertyPublished struct {
	Published string `json:"published,omitempty"`
}

// PropertySummary string
type PropertySummary struct {
	Summary string `json:"summary,omitempty"`
}

// PropertySummaryMap map of lang to string
type PropertySummaryMap struct {
	SummaryMap map[string]string `json:"summaryMap,omitempty"`
}

// PropertyUpdated string as a functional type (single)
type PropertyUpdated struct {
	Updated string `json:"updated,omitempty"`
}

// PropertyUrl inherits Link
type PropertyUrl struct {
	Url interface{} `json:"url,omitempty"`
}

// PropertyTo inherits Object or Link
type PropertyTo struct {
	To interface{} `json:"to,omitempty"`
}

// PropertyMediaType string (mimetype) as a functional type (single)
type PropertyMediaType struct {
	MediaType string `json:"mediaType,omitempty"`
}

// PropertyDuration string as a functional type (single)
type PropertyDuration struct {
	Duration string `json:"duration,omitempty"`
}

// PropertyHref string as a functional type (single)
type PropertyHref struct {
	Href string `json:"href,omitempty"`
}

// PropertyHrefLang string as a functional type (single)
type PropertyHrefLang struct {
	HrefLang string `json:"href,omitempty"`
}

// PropertyRel string as a functional type (single)
type PropertyRel struct {
	Rel []string `json:"rel,omitempty"`
}

// PropertyHeight string as a functional type (single)
type PropertyHeight struct {
	Height uint `json:"height,omitempty"`
}

// PropertyWidth string as a functional type (single)
type PropertyWidth struct {
	Width uint `json:"width,omitempty"`
}

// PropertyTotalItems string as a functional type (single)
type PropertyTotalItems struct {
	TotalItems uint `json:"totalItems,omitempty"`
}

// PropertyStartIndex string as a functional type (single)
type PropertyStartIndex struct {
	StartIndex uint `json:"startIndex,omitempty"`
}
