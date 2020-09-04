package main

import (
	"github.com/bukowa/kube"
	"github.com/bukowa/kube/examples"
	"github.com/bukowa/kube/hooks"
	kv1apps "github.com/bukowa/kube/kubernetes/apps/v1"
	kv1core "github.com/bukowa/kube/kubernetes/core/v1"
	"k8s.io/client-go/kubernetes"
)

const (
	// each should have unique value
	Namespace  kv1core.Namespace  = "namespace"
	Deployment kv1apps.Deployment = "deployment"
)

var (
	Container = kube.NewContainer(Deployment, Namespace)
	Hooks     = kube.Hooks{
		kube.PreCreate: []kube.Hook{
			hooks.SetLabels(map[string]string{"label": "value"}),
		},
		kube.PreGet: []kube.Hook{
			func(c kube.Container) error {
				c.Deployment(Deployment)
				return nil
			},
		},
	}
)

func main() {
	kubeClient := getClient()
	controller := kube.NewController(Container, nil, kube.WithKubernetesClient("namespace", kubeClient))
	controller.RegisterHooks(Hooks)
	controller.CreateContainer()
	controller.ForEachResource(func(res kube.Resource) {
		res.SetName("my-name")
	})
	_, err := controller.CreateKind(Namespace)
	if err != nil {
		panic(err)
	}

}

func getClient() *kubernetes.Clientset {
	kubeClient, err := examples.GetKubernetesClient()
	if err != nil {
		panic(kubeClient)
	}
	return kubeClient
}
