package main

import (
	"fmt"

	"github.com/iancoleman/strcase"
)

func newClassicDashboard(a api) entityType {
	const entity = "ClassicDashboard"
	return entityType{
		API:                  a,
		GoType:               fmt.Sprintf("%s%s", a.GoPrefix, entity),
		SwaggerType:          entity,
		SwaggerModel:         "GrafanaDashboard",
		SwaggerClient:        fmt.Sprintf("%s.%s", a.Client, entity),
		SwaggerClientPackage: strcase.ToSnake(entity),
		DryRun:               true,
		UpdateUnsupported:    false,
	}
}
