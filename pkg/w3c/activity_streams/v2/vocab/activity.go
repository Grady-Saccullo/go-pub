package vocab

type Activity interface {
	GetCreate() ActivityCreate
	SetCreate(create ActivityCreate)

	GetAccept() ActivityAccept
	SetAccept(accept ActivityAccept)

	GetFollow() ActivityFollow
	SetFollow(follow ActivityFollow)

	GetLike() ActivityLike
	SetLike(like ActivityLike)
}
