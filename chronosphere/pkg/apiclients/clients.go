package apiclients

import (
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configunstable"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1"
)

// Clients groups all the generated API clients used by TF.
type Clients struct {
	ConfigUnstable *configunstable.Client
	ConfigV1       *configv1.Client
}
