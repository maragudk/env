package env_test

import (
	"os"
	"testing"
	"time"

	"github.com/maragudk/env"
	"github.com/matryer/is"
)

func TestGetStringOrDefault(t *testing.T) {
	t.Run("gets the string value from the environment", func(t *testing.T) {
		is := is.New(t)
		defer setenv("hat", "party")()
		v := env.GetStringOrDefault("hat", "regular")
		is.Equal("party", v)
	})

	t.Run("gets the default value if not set", func(t *testing.T) {
		is := is.New(t)
		v := env.GetStringOrDefault("hat", "regular")
		is.Equal("regular", v)
	})
}

func TestGetIntOrDefault(t *testing.T) {
	t.Run("gets the int value from the environment", func(t *testing.T) {
		is := is.New(t)
		defer setenv("hats", "2")()
		v := env.GetIntOrDefault("hats", 1)
		is.Equal(2, v)
	})

	t.Run("gets the default value if not set", func(t *testing.T) {
		is := is.New(t)
		v := env.GetIntOrDefault("hats", 1)
		is.Equal(1, v)
	})

	t.Run("gets the default value if not an int", func(t *testing.T) {
		is := is.New(t)
		defer setenv("hats", "notanumber")()
		v := env.GetIntOrDefault("hats", 1)
		is.Equal(1, v)
	})
}

func TestGetBoolOrDefault(t *testing.T) {
	t.Run("gets the bool value from the environment", func(t *testing.T) {
		is := is.New(t)
		defer setenv("hats", "true")()
		v := env.GetBoolOrDefault("hats", true)
		is.Equal(true, v)
	})

	t.Run("gets the default value if not set", func(t *testing.T) {
		is := is.New(t)
		v := env.GetBoolOrDefault("hats", false)
		is.Equal(false, v)
	})

	t.Run("gets the default value if not a bool", func(t *testing.T) {
		is := is.New(t)
		defer setenv("hats", "notabool")()
		v := env.GetBoolOrDefault("hats", false)
		is.Equal(false, v)
	})
}

func TestGetDurationOrDefault(t *testing.T) {
	t.Run("gets the duration value from the environment", func(t *testing.T) {
		is := is.New(t)
		defer setenv("wearhatfor", "1m")()
		v := env.GetDurationOrDefault("wearhatfor", time.Second)
		is.Equal(time.Minute, v)
	})

	t.Run("gets the default value if not set", func(t *testing.T) {
		is := is.New(t)
		v := env.GetDurationOrDefault("wearhatfor", time.Second)
		is.Equal(time.Second, v)
	})

	t.Run("gets the default value if not a bool", func(t *testing.T) {
		is := is.New(t)
		defer setenv("wearhatfor", "notaduration")()
		v := env.GetDurationOrDefault("wearhatfor", time.Second)
		is.Equal(time.Second, v)
	})
}

func setenv(k, v string) func() {
	if err := os.Setenv(k, v); err != nil {
		panic(err)
	}
	return func() {
		if err := os.Unsetenv(k); err != nil {
			panic(err)
		}
	}
}
