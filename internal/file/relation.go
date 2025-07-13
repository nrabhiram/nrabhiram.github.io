package file

type RelationStatus string

const (
	ELDER_SIBLING   RelationStatus = "ELDER SIBLING"
	YOUNGER_SIBLING RelationStatus = "YOUNGER_SIBLING"
	INDETERMINATE   RelationStatus = "INDETERMINATE"
	DISTANT         RelationStatus = "DISTANT"
	PARENT          RelationStatus = "PARENT"
	CHILD           RelationStatus = "CHILD"
)

type Relation struct {
	FirstParty  RelationStatus
	SecondParty RelationStatus
}

func MakeRelation(firstParty RelationStatus, secondParty RelationStatus) Relation {
	return Relation{
		firstParty,
		secondParty,
	}
}
