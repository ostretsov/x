package x

import (
	"os"
	"testing"
)

func TestMustGetenv(t *testing.T) {
	t.Run("panics", func(t *testing.T) {
		var panicked bool
		fn := func() {
			defer func() {
				if r := recover(); r != nil {
					panicked = true
				}
			}()

			MustGetenv("ENVVARDOESNTEXIST111")
		}
		fn()
		if panicked == false {
			t.Errorf("expected panic")
		}
	})

	t.Run("getEnvVar", func(t *testing.T) {
		key := "ENVVARTOSET111"
		value := "VALUE"

		defer func() {
			_ = os.Unsetenv(key)
		}()

		err := os.Setenv(key, value)
		if err != nil {
			t.Error("failed set env var", err)
		}
		got := MustGetenv(key)
		if got != value {
			t.Errorf("want %s; got %s", value, got)
		}
	})
}
