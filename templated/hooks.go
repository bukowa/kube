package templated

import (
	"bytes"
	"github.com/bukowa/kube"
	k8syaml "k8s.io/apimachinery/pkg/util/yaml"
)

// TemplateHooks are registered by default, they execute templates before each step
var TemplateHooks = func(tc *Container, data interface{}) kube.Hooks {
	return kube.Hooks{
		kube.PreGet:    []kube.Hook{TemplateHook(tc, data)},
		kube.PreCreate: []kube.Hook{TemplateHook(tc, data)},
		kube.PreDelete: []kube.Hook{TemplateHook(tc, data)},
	}
}

// TemplateHook executes templates with given data
var TemplateHook = func(tc *Container, data interface{}) kube.Hook {
	return func(c kube.Container) (err error) {
		c.ForEachKind(func(k kube.Kind) {
			buff := bytes.NewBuffer(nil)
			if err = tc.templates.ExecuteTemplate(buff, k.Name(), data); err != nil {
				return
			}
			if err = k8syaml.NewYAMLOrJSONDecoder(buff, buff.Len()).Decode(c.GetResource(k)); err != nil {
				return
			}
		})
		return
	}
}
