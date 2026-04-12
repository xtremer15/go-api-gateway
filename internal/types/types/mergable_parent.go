package types

type MergeableParent interface {
	GetResourceID() int
	AttachChildren(children any, childType string) any
	GetChildren(childrenID int) any
}
