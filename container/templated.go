package container

import (
	"bytes"
	"github.com/bukowa/kube"
	"strings"
	"text/template"
)

// TemplatedContainer implements kube.Container
// using templates to manipulate kube.Resource's
type TemplatedContainer struct {
	kube.Container
	templates *template.Template
}

// NewTemplatesContainer creates TemplatedContainer parsing templates and
// executing them - making sure they won't panic at runtime
// by default provided path is parsed, and if patter matches no files
// we parse relative path `templates/*.yaml` - this handles case
// when TemplatedContainer is used by tests in the package itself
func NewTemplatesContainer(path string, kinds ...kube.Kind) *TemplatedContainer {
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

	return &TemplatedContainer{
		Container: kube.NewContainer(kinds...),
		templates: templates,
	}
}
