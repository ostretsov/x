package x

import "os"

func MustGetenv(key string) string {
	v := os.Getenv(key)
	if v == "" {
		panic("failed to get env var " + key)
	}
	return v
}
