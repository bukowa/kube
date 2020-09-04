package kube

import (
	"github.com/bukowa/kube"
	"github.com/bukowa/kube/kubernetes/core/v1"
	v1core "k8s.io/api/core/v1"
	"testing"
)

const TestSecret v1.Secret = "secret"
const TestService v1.Service = "service"

func TestNewContainer(t *testing.T) {
	container := kube.NewContainer(TestSecret, TestService)
	if container.Secret(TestSecret) == nil {
		t.Error()
	}
	if container.Service(TestService) == nil {
		t.Error()
	}
	if v, ok := container.GetResource(TestSecret).(*v1core.Secret); !ok {
		t.Error(v)
	}
	if v, ok := container.GetResource(TestService).(*v1core.Service); !ok {
		t.Error(v)
	}
	obj := container.GetResource(TestSecret)
	obj.SetName("test-name")
	if container.GetResource(TestSecret).GetName() != "test-name" {
		t.Error()
	}
	obj = &v1core.Service{}
	container.Update(TestService, obj)
	if container.GetResource(TestService) != obj {
		t.Error()
	}
	container.ForEachResource(func(resource kube.Resource) {
		resource.SetName("new-name")
	})
	container.ForEachKind(func(kind kube.Kind) {
		if container.GetResource(kind).GetName() != "new-name" {
			t.Error()
		}
	})
}
