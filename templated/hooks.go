package templated

import (
	"bytes"
	"github.com/bukowa/kube"
	k8syaml "k8s.io/apimachinery/pkg/util/yaml"
)

// TemplateHooks are registered by default, they execute templates before each step
var TemplateHooks = func(tc *Container, data interface{}) kube.Hooks {
	return kube.Hooks{
		kube.PreGet:    {TemplateHook(tc, data)},
		kube.PreCreate: {TemplateHook(tc, data)},
		kube.PreDelete: {TemplateHook(tc, data)},
	}
}

// TemplateHook executes templates with given data
var TemplateHook = func(tc *Container, data interface{}) kube.Hook {
	return func(c kube.Container) (err error) {
		for _, kind := range c.Kinds() {
			if err = ExecuteTemplate(tc, kind, data); err != nil {
				return
			}
		}
		return
	}
}

func ExecuteTemplate(tc *Container, k kube.Kind, d interface{}) (err error) {
	buff := bytes.NewBuffer(nil)
	if err = tc.templates.ExecuteTemplate(buff, k.Name(), d); err != nil {
		return
	}
	// decode into Kind's Resource
	if err = k8syaml.NewYAMLOrJSONDecoder(buff, buff.Len()).Decode(tc.GetResource(k)); err != nil {
		return
	}
	return nil
}
