package vocab

type Activity interface {
	GetActivityCreate() ActivityCreate
	GetActivityAccept() ActivityAccept
	GetActivityFollow() ActivityFollow
}
