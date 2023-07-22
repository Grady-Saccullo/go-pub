package activityPub

import (
	activityStream "github.com/Grady-Saccullo/activity-pub-go/internal/types/w3c/activity_streams/_old_"
)

type Actor struct {
	activityStream.Actor
	PropertyInbox
	PropertyOutbox
	PropertyFollowing
	PropertyFollowers
	PropertyLiked
	PropertyStreams
	PropertyPreferredUsername
	PropertyEndpoints
}
