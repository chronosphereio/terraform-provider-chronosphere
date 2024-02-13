package unstabletest

import (
	"os"
	"testing"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/unstable"
	"github.com/stretchr/testify/require"
)

func TestWithEnabled(t *testing.T) {
	require.False(t, unstable.Enabled())

	WithEnabled(t)

	require.True(t, unstable.Enabled())
	require.Equal(t, os.Getenv(unstable.Env), "1")
}
