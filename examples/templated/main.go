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
	controller := templated.NewController(Container, &data, kube.Hooks{}, kube.WithKubernetesClient(namespace, kubeClient))
	errs := controller.CreateContainer()
	if errs != nil {
		for _, err := range errs {
			log.Print(err)
		}
	}
	deployment := controller.Deployment(Deployment)
	log.Print(deployment.Namespace, deployment.Name, deployment.ObjectMeta.UID, deployment.ObjectMeta.CreationTimestamp)
	// 2020/09/04 19:39:26 templated-namespacetemplated-deployment 000099af-14e4-4c1c-928b-581df9e1b0fa 2020-09-04 19:39:27 +0200 CEST

	deployment = controller.Deployment(Deployment)
	controller.GetContainer()
	log.Print(deployment.Namespace, deployment.Name, deployment.ObjectMeta.UID, deployment.ObjectMeta.CreationTimestamp)

}
