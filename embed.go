//go:build embedexamples
// +build embedexamples

package main

import (
	"embed"
)

//go:embed examples/*
var devFS embed.FS
