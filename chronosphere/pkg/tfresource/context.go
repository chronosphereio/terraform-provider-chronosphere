package tfresource

import (
	"context"
)

type tfResourceContextKey struct{}

// NewContext creates a new context with the TF resource name set
func NewContext(ctx context.Context, resourceName string) context.Context {
	return context.WithValue(ctx, tfResourceContextKey{}, resourceName)
}

// FromContext returns the resource name from the context
func FromContext(ctx context.Context) (string, bool) {
	v, ok := ctx.Value(tfResourceContextKey{}).(string)
	if !ok {
		return "", false
	}
	return v, true
}
