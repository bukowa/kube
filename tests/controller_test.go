package kube

import (
	"context"
	"github.com/bukowa/kube"
	v1meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"testing"
)

func TestNewController(t *testing.T) {
	c, err := GetKubernetesClient()
	if err != nil {
		t.Error(err)
	}
	checkNamespace(c)
	_, err = c.CoreV1().Namespaces().Create(context.TODO(), Namespace, v1meta.CreateOptions{})
	if err != nil {
		t.Error(err)
	}
	defer c.CoreV1().Namespaces().Delete(context.Background(), namespace, v1meta.DeleteOptions{})

	contr := kube.NewController(namespace, kube.NewContainer(TestSecret, TestService), c, kube.OptionCreateOpts(v1meta.CreateOptions{
		DryRun: []string{"All"},
	}))
	if contr.Service(TestService) == nil {
		t.Error()
	}
	errs := contr.CreateContainer()
	if errs != nil {
		t.Error(errs)
	}
}
