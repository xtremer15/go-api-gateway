package interfaces

import "api-gateway/internal/types/types"

type AllowedTypes interface {
	types.User | types.Comment | types.Post
}
