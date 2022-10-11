package x

func Filter[T any](s []T, fn func(T) bool) []T {
	var filtered []T
	for _, v := range s {
		if fn(v) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

func Contains[T comparable](sl []T, v T) bool {
	for _, s := range sl {
		if s == v {
			return true
		}
	}
	return false
}
