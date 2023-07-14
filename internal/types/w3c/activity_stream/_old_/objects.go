package _old_

import (
	"encoding/json"
	"errors"
)

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

type ObjectCoreI interface {
	GetPropertyAttachment() PropertyAttachment
	GetPropertyAttributedTo() PropertyAttributedTo
}

type ObjectCore struct {
	JsonLDContext JsonLDContext `json:"@context,omitempty"`
	ID            PropertyId    `json:"id,omitempty"`
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

func NewObjectFromJson(d []byte) (*Object, error) {
	var o = &Object{}
	if err := o.LoadJson(d); err != nil {
		return nil, err
	}
	return o, nil
}

func NewObject() *Object {
	var o = Object{
		Type: ObjectTypeObject,
	}
	return &o
}

type Object struct {
	ObjectCore
	Type ObjectType `json:"type"`
	data *[]byte
}

func (o *Object) LoadJson(d []byte) error {
	o.data = &d
	return json.Unmarshal(d, &o)
}

func (o *Object) GetExtension(v any) error {
	return json.Unmarshal(*o.data, v)
}

func (o *Object) SetExtension(v any) error {
	return nil
}

func (o *Object) ValidateType(d ...interface{}) error {
	switch o.Type {
	case ObjectTypeObject:
		return nil
	case ObjectTypeArticle:
		return nil
	case ObjectTypeAudio:
		return nil
	case ObjectTypeDocument:
		return nil
	case ObjectTypeEvent:
		return nil
	case ObjectTypeImage:
		return nil
	case ObjectTypeNote:
		return nil
	case ObjectTypePage:
		return nil
	case ObjectTypePlace:
		return nil
	case ObjectTypeProfile:
		return nil
	case ObjectTypeRelationship:
		return nil
	case ObjectTypeTombstone:
		return nil
	case ObjectTypeVideo:
		return nil
	default:
		for _, v := range d {
			if v == string(o.Type) {
				return nil
			}
		}
	}
	return errors.New("unknown object type")
}

func NewRelationshipFromJson(d []byte) (*Relationship, error) {
	var o = Relationship{}
	if err := o.LoadJson(d); err != nil {
		return nil, err
	}
	return &o, nil
}

type Relationship struct {
	Object
	PropertyObject
	PropertyRelationship
	PropertySubject
}

func NewArticleFromJson(d []byte) (*Article, error) {
	var o = Article{}
	if err := o.LoadJson(d); err != nil {
		return nil, err
	}
	return &o, nil
}

type Article struct {
	Object
}

func NewDocumentFromJson(d []byte) (*Document, error) {
	var o = Document{}
	if err := o.LoadJson(d); err != nil {
		return nil, err
	}
	return &o, nil
}

type Document struct {
	Object
}

func NewAudioFromJson(d []byte) (*Audio, error) {
	var o = Audio{}
	if err := o.LoadJson(d); err != nil {
		return nil, err
	}
	return &o, nil
}

type Audio struct {
	Document
}

func NewImageFromJson(d []byte) (*Image, error) {
	var o = Image{}
	if err := o.LoadJson(d); err != nil {
		return nil, err
	}
	return &o, nil
}

type Image struct {
	Document
}

func NewVideoFromJson(d []byte) (*Video, error) {
	var o = Video{}
	if err := o.LoadJson(d); err != nil {
		return nil, err
	}
	return &o, nil
}

type Video struct {
	Document
}

func NewNoteFromJson(d []byte) (*Note, error) {
	var o = Note{}
	if err := o.LoadJson(d); err != nil {
		return nil, err
	}
	return &o, nil
}

type Note struct {
	Object
}

func NewPageFromJson(d []byte) (*Page, error) {
	var o = Page{}
	if err := o.LoadJson(d); err != nil {
		return nil, err
	}
	return &o, nil
}

type Page struct {
	Document
}

func NewEventFromJson(d []byte) (*Event, error) {
	var o = Event{}
	if err := o.LoadJson(d); err != nil {
		return nil, err
	}
	return &o, nil
}

type Event struct {
	Object
}

func NewPlaceFromJson(d []byte) (*Place, error) {
	var o = Place{}
	if err := o.LoadJson(d); err != nil {
		return nil, err
	}
	return &o, nil
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

func NewProfileFromJson(d []byte) (*Profile, error) {
	var o = Profile{}
	if err := o.LoadJson(d); err != nil {
		return nil, err
	}
	return &o, nil
}

type Profile struct {
	Object
	PropertyDescribes
}

func NewTombstoneFromJson(d []byte) (*Tombstone, error) {
	var o = Tombstone{}
	if err := o.LoadJson(d); err != nil {
		return nil, err
	}
	return &o, nil
}

type Tombstone struct {
	Object
	PropertyDeleted
	PropertyFormerType
}
