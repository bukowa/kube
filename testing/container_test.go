package testing

import (
	"github.com/bukowa/kube"
	kv1apps "github.com/bukowa/kube/kubernetes/apps/v1"
	kv1core "github.com/bukowa/kube/kubernetes/core/v1"
	kv1beta1net "github.com/bukowa/kube/kubernetes/networking/v1beta1"
	"testing"
)

const (
	deployment            kv1apps.Deployment            = "deployment"
	configmap             kv1core.ConfigMap             = "configmap"
	namespace             kv1core.Namespace             = "namespace"
	persistentvolumeclaim kv1core.PersistentVolumeClaim = "persistentvolumeclaim"
	secret                kv1core.Secret                = "secret"
	service               kv1core.Service               = "service"
	ingress               kv1beta1net.Ingress           = "ingress"
)

var (
	group = []kube.Kind{
		namespace, secret, configmap, persistentvolumeclaim, deployment, service, ingress,
	}
)

// it doesnt panic with proper kinds
func TestNewContainer(t *testing.T) {
	kube.NewContainer(group...)
}

// duplicate kube.Kind are not allowed
func TestNewContainerDuplicates(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("did not panic")
		} else {
			if r != "duplicate kind secret" {
				t.Error(r)
			}
		}
	}()
	kube.NewContainer(secret, secret)
}

// we should return a copy
func TestNewContainerCopy(t *testing.T) {
	c := kube.NewContainer(group...)
	c2 := c.Copy()

	c.GetResource(secret).SetName("1")
	c2.GetResource(secret).SetName("2")

	if c.GetResource(secret).GetName() != "1" {
		t.Error()
	}

	if c2.GetResource(secret).GetName() != "2" {
		t.Error()
	}

	if c.GetResource(secret) == c2.GetResource(secret) {
		t.Error()
	}
}
