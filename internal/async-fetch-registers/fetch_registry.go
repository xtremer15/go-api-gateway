package registry

type Registry map[string]func() <-chan any

var GlobalRegistry = make(Registry)

func RegisterAsyncFn(key string, fn func() <-chan any) {
	GlobalRegistry[key] = fn
	// if strings.Contains(queryParams, key) {
	// }
}

func GetRegisteredFn(key string) func() <-chan any {
	return GlobalRegistry[key]
}
