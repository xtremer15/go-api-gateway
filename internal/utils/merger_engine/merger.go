package merger

import (
	"api-gateway/internal/types/types"
)

func Merger(parents []types.MergeableParent, children []types.MergeableChild) {

	childrenByParent := make(map[int][]types.MergeableChild)

	for _, child := range children {
		parentID := child.GetParentID()
		childrenByParent[parentID] = append(childrenByParent[parentID], child)
	}

	for _, parent := range parents {
		parentID := parent.GetResourceID()
		if matchingComments, exists := childrenByParent[parentID]; exists {
			childType := matchingComments[0].GetChildType()
			parent.AttachChildren(matchingComments, childType)
		}
	}

}
