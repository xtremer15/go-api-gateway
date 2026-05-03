package utils

import "api-gateway/internal/types/types"

func toParent[T any](data []T) []types.MergeableParent {
	var result []types.MergeableParent
	for i := range data {
		if p, ok := any(&data[i]).(types.MergeableParent); ok {
			result = append(result, p)
		}
	}
	return result
}

func toChildren[T any](data []T) []types.MergeableChild {
	var result []types.MergeableChild
	for i := range data {
		if c, ok := any(&data[i]).(types.MergeableChild); ok {
			result = append(result, c)
		}
	}
	return result
}

func TypeConvertor[T any](data []T, convertTo string) ([]types.MergeableParent, []types.MergeableChild) {
	switch convertTo {
	case "parent":
		return toParent(data), nil
	case "children":
		return nil, toChildren(data)
	default:
		return nil, nil
	}
}
