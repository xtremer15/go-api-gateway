package utils

import "api-gateway/internal/types/types"

func toParent[T types.MergeableParent](data []T) []types.MergeableParent {
	var result []types.MergeableParent
	for i := range data {
		result = append(result, data[i]) // Note: you might need to pass pointers here depending on how you call it!
	}
	return result
}

func toChildren[T types.MergeableChild](data []T) []types.MergeableChild {
	var result []types.MergeableChild
	for i := range data {
		result = append(result, data[i]) // Note: you might need to pass pointers here depending on how you call it!
	}
	return result
}

func TypeConvertor[T types.MergeableParent, U types.MergeableChild](data []T, convertTo string) ([]types.MergeableParent, []types.MergeableChild) {
	switch convertTo {
	case "parent":
		return toParent[types.MergeableParent](data), nil
	case "children":
		return nil, toChildren[types.MergeableChild](data)
	default:
		return nil, nil
	}
}
