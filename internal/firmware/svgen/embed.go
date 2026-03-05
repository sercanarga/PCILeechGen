package svgen

import "embed"

//go:embed templates/*.sv.tmpl
var templateFS embed.FS

// mustReadTemplate reads an embedded template file.
func mustReadTemplate(name string) string {
	data, err := templateFS.ReadFile("templates/" + name)
	if err != nil {
		panic("embedded template not found: " + name + ": " + err.Error())
	}
	return string(data)
}
