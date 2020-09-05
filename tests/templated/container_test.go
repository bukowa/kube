package templated

import (
	"github.com/bukowa/kube"
	kv1apps "github.com/bukowa/kube/kubernetes/apps/v1"
	kv1core "github.com/bukowa/kube/kubernetes/core/v1"
	kv1beta1net "github.com/bukowa/kube/kubernetes/networking/v1beta1"
	"github.com/bukowa/kube/templated"
	"testing"
)

const (
	Deployment            kv1apps.Deployment            = "deployment.yaml"
	Secret                kv1core.Secret                = "secret.yaml"
	Configmap             kv1core.ConfigMap             = "configmap.yaml"
	Namespace             kv1core.Namespace             = "namespace.yaml"
	Persistentvolumeclaim kv1core.PersistentVolumeClaim = "persistentvolumeclaim.yaml"
	Service               kv1core.Service               = "service.yaml"
	Ingress               kv1beta1net.Ingress           = "ingress.yaml"
)

var Group = []kube.Kind{
	Deployment, Secret, Configmap, Namespace, Persistentvolumeclaim, Service, Ingress,
}

type Data struct {
	Name string
}

func TestContainer(t *testing.T) {
	c := templated.NewContainer("tests/templated/templates", Group...)
	d := &Data{Name: "new"}

	if err := templated.TemplateHook(c, d)(c); err != nil {
		t.Error(err)
	}

	for _, kind := range Group {
		if resource := c.GetResource(kind); resource == nil {
			t.Errorf("resource: %s is nil", kind.Name())
		} else if resource.GetName() != "new" {
			t.Errorf("resource: %s name is wrong", kind.Name())
		}
	}
}
