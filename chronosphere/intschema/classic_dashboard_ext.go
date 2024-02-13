package intschema

import (
	"io"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/hclmarshal"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfid"
)

// ClassicDashboard is an alias to the grafana dashboard type. Due to
// registry/intschema limitations, we simply wrap the grafana intschema type
// and override marshalling bits to use the new resource name.
type ClassicDashboard GrafanaDashboard

func (o *ClassicDashboard) MarshalHCL(w io.Writer) error {
	m := hclmarshal.New()
	b := m.AddResource("chronosphere_classic_dashboard", o.HCLID)
	if err := hclmarshal.MarshalIntSchema(o, b); err != nil {
		return err
	}
	return m.MarshalTo(w)
}

func (o *ClassicDashboard) Ref() tfid.ID {
	if o.HCLID == "" {
		panic("Ref is only valid when schema structs are used for marshalling")
	}

	return tfid.Ref{
		Type: "chronosphere_classic_dashboard",
		ID:   o.HCLID,
	}.AsID()
}
