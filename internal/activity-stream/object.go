package activityStream

type ObjectType string

const (
	ObjectTypeObject       ObjectType = "Object"
	ObjectTypeArticle      ObjectType = "Article"
	ObjectTypeAudio        ObjectType = "Audio"
	ObjectTypeDocument     ObjectType = "Document"
	ObjectTypeEvent        ObjectType = "Event"
	ObjectTypeImage        ObjectType = "Image"
	ObjectTypeNote         ObjectType = "Note"
	ObjectTypePage         ObjectType = "Page"
	ObjectTypePlace        ObjectType = "Place"
	ObjectTypeProfile      ObjectType = "Profile"
	ObjectTypeRelationship ObjectType = "Relationship"
	ObjectTypeTombstone    ObjectType = "Tombstone"
	ObjectTypeVideo        ObjectType = "Video"
)

type ObjectCore struct {
	ID PropertyId `json:"id,omitempty"`
	PropertyAttachment
	PropertyAttributedTo
	PropertyAudience
	PropertyBcc
	PropertyBto
	PropertyCc
	PropertyContent
	PropertyContentMap
	PropertyContext
	PropertyDuration
	PropertyGenerator
	PropertyIcon
	PropertyImage
	PropertyInReplyTo
	PropertyLocation
	PropertyMediaType
	PropertyName
	PropertyNameMap
	PropertyPreview
	PropertyPublished
	PropertyReplies
	PropertyStartTime
	PropertySummary
	PropertySummaryMap
	PropertyEndTime
	PropertyUpdated
	PropertyUrl
	PropertyTag
	PropertyTo
}

type Object struct {
	JsonLDContext JsonLDContext `json:"@context"`
	Type          ObjectType    `json:"type"`
	ObjectCore
}

type Relationship struct {
	Object
	PropertyObject
	PropertyRelationship
	PropertySubject
}

type Article struct {
	Object
}

type Document struct {
	Object
}

type Audio struct {
	Document
}

type Image struct {
	Document
}

type Video struct {
	Document
}

type Note struct {
	Object
}

type Page struct {
	Document
}

type Event struct {
	Object
}

type Place struct {
	Object
	PropertyAccuracy
	PropertyAltitude
	PropertyLatitude
	PropertyLongitude
	PropertyRadius
	PropertyUnits
}

type Profile struct {
	Object
	PropertyDescribes
}

type Tombstone struct {
	Object
	PropertyDeleted
	PropertyFormerType
}
