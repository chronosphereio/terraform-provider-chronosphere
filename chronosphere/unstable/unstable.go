package unstable

import (
	"context"
	"errors"
	"sync"

	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Env is the unstable environment variable.
const Env = "UNSTABLE"

var (
	_mu          sync.RWMutex
	_initialized = false
	_enabled     = false
)

// Enabled returns true if the unstable code path is enabled. Panics if Set has
// not been called yet.
func Enabled() bool {
	v, err := SafeEnabled()
	if err != nil {
		panic(err)
	}
	return v
}

// SafeEnabled returns true if the unstable code path is enabled. Returns error
// if Set has not been called yet.
func SafeEnabled() (bool, error) {
	_mu.RLock()
	defer _mu.RUnlock()

	if !_initialized {
		return false, errors.New("unstable.Set must be called before unstable.Enabled")
	}
	return _enabled, nil
}

// Set sets the unstable code path to either enabled or disabled.
func Set(ctx context.Context, enabled bool) {
	_mu.Lock()
	defer _mu.Unlock()

	if _initialized {
		tflog.Warn(ctx, "unstable.Set called more than once", map[string]any{
			"old_val": _enabled,
			"new_val": enabled,
		})
	}
	if enabled {
		tflog.Warn(ctx, "unstable code path enabled")
	} else {
		tflog.Info(ctx, "unstable code path disabled")
	}
	_initialized = true
	_enabled = enabled
}
