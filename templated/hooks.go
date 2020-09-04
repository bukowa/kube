package templated

import (
	"bytes"
	"github.com/bukowa/kube"
	k8syaml "k8s.io/apimachinery/pkg/util/yaml"
)

// Hooks are registered by default, they execute templates before each step
var Hooks = func(tc *Container, data interface{}) kube.Hooks {
	return kube.Hooks{
		kube.PreGet:    []kube.Hook{Hook(tc, data)},
		kube.PreCreate: []kube.Hook{Hook(tc, data)},
		kube.PreDelete: []kube.Hook{Hook(tc, data)},
	}
}

// Hook executes templates with given data
var Hook = func(tc *Container, data interface{}) kube.Hook {
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
