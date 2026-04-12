package types

type MergeableChild interface {
	GetParentID() int
	GetChildKey() int
	GetChildType() string
}
