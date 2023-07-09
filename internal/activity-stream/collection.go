package activityStream

type Collection struct {
	ObjectCore
	PropertyTotalItems
	PropertyCurrent
	PropertyFirst
	PropertyLast
	PropertyItems
}

type OrderedCollection struct {
	Collection
	PropertyOrderedItems
}

type CollectionPage struct {
	Collection
	PropertyNext
	PropertyPrev
	PropertyPartOf
}

type OrderedCollectionPage struct {
	CollectionPage
	OrderedCollection
	PropertyStartIndex
}
