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

package configv1

import (
	"github.com/go-openapi/strfmt"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/client"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/transport"
)

type Client = client.ConfigV1API

func NewClient(component transport.Component, org, apiToken, entityNamespace string) (*Client, error) {
	t, err := transport.New(transport.Params{
		Component:       component,
		Org:             org,
		APIToken:        apiToken,
		EntityNamespace: entityNamespace,
	})
	if err != nil {
		return nil, err
	}
	return client.New(t, strfmt.Default), nil
}
