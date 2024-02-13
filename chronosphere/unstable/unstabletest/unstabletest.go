package unstabletest

import (
	"context"
	"os"
	"testing"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/unstable"
)

func init() {
	unstable.Set(context.Background(), os.Getenv(unstable.Env) == "1")
}

// WithEnabled runs a test with unstable enabled. It also sets the unstable
// environment variable to enabled, such that all subprocesses initialized by
// the test also have unstable enabled.
//
// At the end of the test, the original values for both the local unstable
// package and the environment variable are restored.
func WithEnabled(t *testing.T) {
	t.Cleanup(withLocal(true))
	t.Cleanup(withEnv(true))
}

func withLocal(enabled bool) (cleanup func()) {
	original, err := unstable.SafeEnabled()
	if err != nil {
		// Treat uninitialized as just false.
		original = false
	}

	unstable.Set(context.Background(), enabled)

	return func() {
		unstable.Set(context.Background(), original)
	}
}

func withEnv(enabled bool) (cleanup func()) {
	original, ok := os.LookupEnv(unstable.Env)

	if enabled {
		os.Setenv(unstable.Env, "1")
	} else {
		os.Unsetenv(unstable.Env)
	}

	return func() {
		if !ok {
			os.Unsetenv(unstable.Env)
		} else {
			os.Setenv(unstable.Env, original)
		}
	}
}
