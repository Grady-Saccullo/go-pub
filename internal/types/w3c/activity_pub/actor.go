package activityPub

type Actor struct {
	PropertyInbox
	PropertyOutbox
	PropertyFollowing
	PropertyFollowers
	PropertyLiked
	PropertyStreams
	PropertyPreferredUsername
	PropertyEndpoints
}
