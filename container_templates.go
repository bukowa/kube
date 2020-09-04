package kube

import (
	"bytes"
	"strings"
	"text/template"
)

// TemplatesContainer implements Container
// using templates to manipulate Resource's
type TemplatesContainer struct {
	Container
	templates *template.Template
}

// NewTemplatesContainer creates TemplatesContainer parsing templates and
// executing them - making sure they won't panic at runtime
// by default provided path is parsed, and if patter matches no files
// we parse relative path `templates/*.yaml` - this handles case
// when TemplatesContainer is used by tests in the package itself
func NewTemplatesContainer(path string, kinds ...Kind) *TemplatesContainer {
	var templates *template.Template

	// parse provided path
	templates, err := template.ParseGlob(path)
	if err != nil && !strings.Contains(err.Error(), "template: pattern matches no files") {
		panic(err)
	} else if err != nil {
		// allow relative path if nothing was found
		if templates, err = template.ParseGlob("templates/*.yaml"); err != nil {
			panic(err)
		}
	}

	// executing templates should not panic
	for _, k := range kinds {
		err := templates.ExecuteTemplate(bytes.NewBuffer(nil), k.Name(), nil)
		if err != nil {
			panic(err)
		}
	}

	return &TemplatesContainer{
		Container: NewContainer(kinds...),
		templates: templates,
	}
}
