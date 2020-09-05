package templated

import (
	"github.com/bukowa/kube"
	kv1apps "github.com/bukowa/kube/kubernetes/apps/v1"
	kv1core "github.com/bukowa/kube/kubernetes/core/v1"
	kv1beta1net "github.com/bukowa/kube/kubernetes/networking/v1beta1"
	"github.com/bukowa/kube/templated"
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

var (
	Container = templated.NewContainer("testing/templated/templates/*.yaml", Group...)
	Group     = []kube.Kind{
		Deployment, Secret, Configmap, Namespace, Persistentvolumeclaim, Service, Ingress,
	}
	Data = &data{Name: "new"}
)

type data struct {
	Name string
}
