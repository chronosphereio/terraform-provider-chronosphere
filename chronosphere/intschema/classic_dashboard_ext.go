// Copyright 2024 Chronosphere Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
