package templated

import (
	"bytes"
	"fmt"
	"github.com/bukowa/kube"
	"os"
	"strings"
	"text/template"
)

// Container implements kube.Container
// using templates to manipulate kube.Resource's
type Container struct {
	kube.Container
	path      string
	templates *template.Template
}

// NewContainer creates Container parsing templates and
// executing them - making sure they won't panic at runtime
// by default provided path is parsed, and if patter matches no files
// we parse relative path `templates/*.yaml` - this handles case
// when Container is used by tests in the package itself
func NewContainer(path string, kinds ...kube.Kind) *Container {
	var templates *template.Template
	// parse provided path
	templates, err := template.ParseGlob(path)
	if err != nil && !strings.Contains(err.Error(), "template: pattern matches no files") {
		panic(err)
	} else if err != nil {
		// allow relative path if nothing was found
		if templates, err = template.ParseGlob("templates/*.yaml"); err != nil {
			panic(fmt.Sprintf("%s | working dir: %s | parsed path: %s", err, wd(), path))
		}
	}

	// executing templates should not panic
	for _, k := range kinds {
		err := templates.ExecuteTemplate(bytes.NewBuffer(nil), k.Name(), nil)
		if err != nil {
			panic(err)
		}
	}

	return &Container{
		Container: kube.NewContainer(kinds...),
		path:      path,
		templates: templates,
	}
}

func (tc *Container) Copy() kube.Container {
	return NewContainer(tc.path, tc.Kinds()...)
}

func (tc *Container) Self() kube.Container {
	return tc
}

func (tc *Container) Templates() *template.Template {
	return tc.templates
}

func wd() string {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return wd
}
