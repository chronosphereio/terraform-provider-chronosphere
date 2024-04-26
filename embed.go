//go:build embedexamples
// +build embedexamples

package main

import (
	"embed"
)

//go:embed examples/*
var examplesFS embed.FS

func Examples() embed.FS {
	return examplesFS
}
