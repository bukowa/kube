package main

import (
	"github.com/bukowa/kube"
	"github.com/bukowa/kube/examples"
	kv1apps "github.com/bukowa/kube/kubernetes/apps/v1"
	kv1core "github.com/bukowa/kube/kubernetes/core/v1"
	"github.com/bukowa/kube/templated"
	"log"
)

var (
	Namespace  kv1core.Namespace  = "namespace.yaml"
	Deployment kv1apps.Deployment = "deployment.yaml"
	Container                     = templated.NewContainer("examples/templated/templates", Namespace, Deployment)
)

type Data struct {
	Namespace string
	Name      string
	Label     string
}

func main() {
	namespace := "templated-namespace"
	data := &Data{
		Namespace: "templated-namespace",
		Name:      "templated-deployment",
		Label:     "hello-world",
	}
	kubeClient, err := examples.GetKubernetesClient()
	if err != nil {
		panic(kubeClient)
	}
	controller := templated.NewController(Container, data, nil, kube.WithKubernetesClient(namespace, kubeClient))
	errs := controller.CreateContainer()
	if errs != nil {
		for _, err := range errs {
			log.Print(err)
		}
	}
	//deployment := controller.Deployment(Deployment)
	//log.Print(deployment.Namespace, deployment.Name, deployment.Spec.Template.Spec.Containers)
}
