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
