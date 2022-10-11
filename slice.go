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
