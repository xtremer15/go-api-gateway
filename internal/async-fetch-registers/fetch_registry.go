package registry

import "strings"

type Registry map[string]func() <-chan any

var GlobalRegistry = make(Registry)

func RegisterAsyncFn(key string, fn func() <-chan any) {
	if strings.Contains(queryParams, key) {
		GlobalRegistry[key] = fn
	}
}

func GetRegisteredFn(key string) func() <-chan any {
	return GlobalRegistry[key]
}
