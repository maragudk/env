package env_test

import (
	"os"
	"testing"
	"time"

	"github.com/maragudk/is"

	"github.com/maragudk/env"
)

func TestGetStringOrDefault(t *testing.T) {
	t.Run("gets the string value from the environment", func(t *testing.T) {
		defer setenv("hat", "party")()
		v := env.GetStringOrDefault("hat", "regular")
		is.Equal(t, "party", v)
	})

	t.Run("gets the default value if not set", func(t *testing.T) {
		v := env.GetStringOrDefault("hat", "regular")
		is.Equal(t, "regular", v)
	})
}

func TestGetIntOrDefault(t *testing.T) {
	t.Run("gets the int value from the environment", func(t *testing.T) {
		defer setenv("hats", "2")()
		v := env.GetIntOrDefault("hats", 1)
		is.Equal(t, 2, v)
	})

	t.Run("gets the default value if not set", func(t *testing.T) {
		v := env.GetIntOrDefault("hats", 1)
		is.Equal(t, 1, v)
	})

	t.Run("gets the default value if not an int", func(t *testing.T) {
		defer setenv("hats", "notanumber")()
		v := env.GetIntOrDefault("hats", 1)
		is.Equal(t, 1, v)
	})
}

func TestGetBoolOrDefault(t *testing.T) {
	t.Run("gets the bool value from the environment", func(t *testing.T) {
		defer setenv("hats", "true")()
		v := env.GetBoolOrDefault("hats", true)
		is.Equal(t, true, v)
	})

	t.Run("gets the default value if not set", func(t *testing.T) {
		v := env.GetBoolOrDefault("hats", false)
		is.Equal(t, false, v)
	})

	t.Run("gets the default value if not a bool", func(t *testing.T) {
		defer setenv("hats", "notabool")()
		v := env.GetBoolOrDefault("hats", false)
		is.Equal(t, false, v)
	})
}

func TestGetDurationOrDefault(t *testing.T) {
	t.Run("gets the duration value from the environment", func(t *testing.T) {
		defer setenv("wearhatfor", "1m")()
		v := env.GetDurationOrDefault("wearhatfor", time.Second)
		is.Equal(t, time.Minute, v)
	})

	t.Run("gets the default value if not set", func(t *testing.T) {
		v := env.GetDurationOrDefault("wearhatfor", time.Second)
		is.Equal(t, time.Second, v)
	})

	t.Run("gets the default value if not a bool", func(t *testing.T) {
		defer setenv("wearhatfor", "notaduration")()
		v := env.GetDurationOrDefault("wearhatfor", time.Second)
		is.Equal(t, time.Second, v)
	})
}

func TestLoad(t *testing.T) {
	t.Run("loads an environment file", func(t *testing.T) {
		defer unsetenv("hat", "hats", "equals")
		err := env.Load("testdata/env")
		is.NotError(t, err)
		hat := env.GetStringOrDefault("hat", "regular")
		is.Equal(t, "party", hat)
		hats := env.GetIntOrDefault("hats", 1)
		is.Equal(t, 2, hats)
	})

	t.Run("loads multiple environment files, and later files take precedence", func(t *testing.T) {
		defer unsetenv("hat", "hats", "equals")
		err := env.Load("testdata/env", "testdata/env2")
		is.NotError(t, err)
		hat := env.GetStringOrDefault("hat", "regular")
		is.Equal(t, "party", hat)
		hats := env.GetIntOrDefault("hats", 1)
		is.Equal(t, 3, hats)
	})

	t.Run("ignores blank lines", func(t *testing.T) {
		defer unsetenv("hat", "hats")
		err := env.Load("testdata/blank")
		is.NotError(t, err)
		hat := env.GetStringOrDefault("hat", "regular")
		is.Equal(t, "party", hat)
	})

	t.Run("errors on bad file", func(t *testing.T) {
		err := env.Load("testdata/invalid")
		is.True(t, err != nil)
		is.Equal(t, "missing equal sign on line 1 in testdata/invalid", err.Error())
	})

	t.Run("ignores comments in environment file", func(t *testing.T) {
		defer unsetenv("hat")
		err := env.Load("testdata/comments")
		is.NotError(t, err)
		hat := env.GetStringOrDefault("hat", "regular")
		is.Equal(t, "party", hat)
	})

	t.Run("gets the string value including equal signs", func(t *testing.T) {
		defer unsetenv("hat", "hats", "equals")
		err := env.Load("testdata/env")
		is.NotError(t, err)
		equals := env.GetStringOrDefault("equals", "")
		is.Equal(t, "somethingwithequalsafter=", equals)
	})
}

func TestMustLoad(t *testing.T) {
	t.Run("loads an environment file", func(t *testing.T) {
		defer unsetenv("hat", "hats", "equals")
		env.MustLoad("testdata/env")
		hat := env.GetStringOrDefault("hat", "regular")
		is.Equal(t, "party", hat)
		hats := env.GetIntOrDefault("hats", 1)
		is.Equal(t, 2, hats)
	})

	t.Run("panics on no such file", func(t *testing.T) {
		recovered := false
		defer func() {
			if err := recover(); err != nil {
				recovered = true
			}
			is.True(t, recovered)
		}()
		env.MustLoad()
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

func unsetenv(ks ...string) {
	for _, k := range ks {
		if err := os.Unsetenv(k); err != nil {
			panic(err)
		}
	}
}
