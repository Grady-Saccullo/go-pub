package activityStream

type ActivityType string

const (
	ActivityTypeActivity        ActivityType = "Activity"
	ActivityTypeAccept          ActivityType = "Accept"
	ActivityTypeAdd             ActivityType = "Add"
	ActivityTypeAnnounce        ActivityType = "Announce"
	ActivityTypeArrive          ActivityType = "Arrive"
	ActivityTypeBlock           ActivityType = "Block"
	ActivityTypeCreate          ActivityType = "Create"
	ActivityTypeDelete          ActivityType = "Delete"
	ActivityTypeDislike         ActivityType = "Dislike"
	ActivityTypeFlag            ActivityType = "Flag"
	ActivityTypeFollow          ActivityType = "Follow"
	ActivityTypeIgnore          ActivityType = "Ignore"
	ActivityTypeInvite          ActivityType = "Invite"
	ActivityTypeJoin            ActivityType = "Join"
	ActivityTypeLeave           ActivityType = "Leave"
	ActivityTypeLike            ActivityType = "Like"
	ActivityTypeListen          ActivityType = "Listen"
	ActivityTypeMove            ActivityType = "Move"
	ActivityTypeOffer           ActivityType = "Offer"
	ActivityTypeQuestion        ActivityType = "Question"
	ActivityTypeReject          ActivityType = "Reject"
	ActivityTypeRead            ActivityType = "Read"
	ActivityTypeRemove          ActivityType = "Remove"
	ActivityTypeTentativeReject ActivityType = "TentativeReject"
	ActivityTypeTentativeAccept ActivityType = "TentativeAccept"
	ActivityTypeTravel          ActivityType = "Travel"
	ActivityTypeUndo            ActivityType = "Undo"
	ActivityTypeUpdate          ActivityType = "Update"
	ActivityTypeView            ActivityType = "View"
)

type ActivityCore struct {
	JsonLDContext JsonLDContext `json:"@context"`
	Type          ActivityType  `json:"type"`
	PropertyActor
	PropertyObject
	PropertyOrigin
	PropertyResult
	PropertyInstrument
	PropertyTarget
}

type Activity struct {
	ObjectCore
	ActivityCore
}

type IntransitiveActivity struct {
	ActivityCore
}

type Accept struct {
	Activity
}

type TentativeAccept struct {
	Accept
}

type Add struct {
	Activity
}

type Arrive struct {
	IntransitiveActivity
}

type Create struct {
	Activity
}

type Delete struct {
	Activity
}

type Follow struct {
	Activity
}

type Ignore struct {
	Activity
}

type Join struct {
	Activity
}

type Leave struct {
	Activity
}

type Like struct {
	Activity
}

type Offer struct {
	Activity
}

type Invite struct {
	Offer
}

type Reject struct {
	Activity
}

type TentativeReject struct {
	Reject
}

type Remove struct {
	Activity
}

type Undo struct {
	Activity
}

type Update struct {
	Activity
}

type View struct {
	Activity
}

type Listen struct {
	Activity
}

type Read struct {
	Activity
}

type Move struct {
	Activity
}

type Travel struct {
	IntransitiveActivity
}

type Announce struct {
	Activity
}

type Block struct {
	Ignore
}

type Flag struct {
	Activity
}

type Dislike struct {
	Activity
}

type Question struct {
	IntransitiveActivity
	PropertyOneOf
	PropertyAnyOf
}
