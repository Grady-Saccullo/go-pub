package activity_streams_v2_vocab

type Activity interface {
	GetActivityCreate() ActivityCreate
	GetActivityAccept() ActivityAccept
	GetActivityFollow() ActivityFollow
	GetActivityLike() ActivityLike
}
