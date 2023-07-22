package activityPub

// PropertyInbox inherits activity_streams.OrderedCollection
type PropertyInbox struct {
	Inbox interface{} `json:"inbox"`
}

// PropertyOutbox inherits activity_streams.OrderedCollection
type PropertyOutbox struct {
	Outbox interface{} `json:"outbox"`
}

// PropertyFollowing inherits activity_streams.Collection or activity_streams.OrderedCollection
type PropertyFollowing struct {
	Following interface{} `json:"following,omitempty"`
}

// PropertyFollowers inherits activity_streams.Collection or activity_streams.OrderedCollection
type PropertyFollowers struct {
	Followers interface{} `json:"followers,omitempty"`
}

// PropertyLiked inherits activity_streams.Collection
type PropertyLiked struct {
	Liked interface{} `json:"liked,omitempty"`
}

// PropertyStreams inherits activity_streams.Collection
type PropertyStreams struct {
	Streams interface{} `json:"liked,omitempty"`
}

// PropertyPreferredUsername inherits activity_streams.Collection
type PropertyPreferredUsername struct {
	PreferredUsername string `json:"liked,omitempty"`
}

type PropertyEndpoints struct {
	Endpoints Endpoints `json:"endpoints"`
}

type Endpoints struct {
	PropertyProxyUrl
	PropertyOauthAuthorizationEndpoint
	PropertyOauthTokenEndpoint
	PropertyProvideClientKey
	PropertySignClientKey
	PropertySharedInbox
}

// PropertyProxyUrl string
type PropertyProxyUrl struct {
	ProxyUrl string `json:"proxyUrl,omitempty"`
}

// PropertyOauthAuthorizationEndpoint string
type PropertyOauthAuthorizationEndpoint struct {
	OauthAuthorizationEndpoint string `json:"oauthAuthorizationEndpoint,omitempty"`
}

// PropertyOauthTokenEndpoint string
type PropertyOauthTokenEndpoint struct {
	OauthTokenEndpoint string `json:"oauthTokenEndpoint,omitempty"`
}

// PropertyProvideClientKey string
type PropertyProvideClientKey struct {
	ProvideClientKey string `json:"provideClientKey,omitempty"`
}

// PropertySignClientKey string
type PropertySignClientKey struct {
	SignClientKey string `json:"signClientKey,omitempty"`
}

// PropertySharedInbox inherits activity_streams.OrderedCollection
type PropertySharedInbox struct {
	SharedInbox string `json:"sharedInbox,omitempty"`
}
