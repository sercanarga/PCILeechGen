package tclgen

import "embed"

//go:embed templates/*.tcl.tmpl
var templateFS embed.FS
