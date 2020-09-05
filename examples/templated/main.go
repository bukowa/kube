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

var (
	namespace = "templated-namespace4"
	data      = &Data{
		Namespace: namespace,
		Name:      "templated-deployment",
		Label:     "hello-world",
	}
)

func main() {
	// create kubernetes client
	kubeClient, err := examples.GetKubernetesClient()
	if err != nil {
		panic(err)
	}
	// create controller for container
	controller := templated.NewController(Container, &data,
		kube.WithKubernetesClient(namespace, kubeClient))

	// after calling get delete deployment
	controller.Configure(kube.WithHooks(map[kube.HookType][]kube.Hook{
		kube.PostGet: {
			func(container kube.Container) error {
				return controller.DeleteKind(Deployment)
			},
		},
	}))

	// create all kinds in container
	errs := controller.CreateContainer()
	if errs != nil {
		for _, err := range errs {
			log.Print(err)
		}
	}

	// grab underlying Deployment
	deployment := controller.Deployment(Deployment)
	log.Print(deployment.Namespace, deployment.Name, deployment.ObjectMeta.UID, deployment.ObjectMeta.CreationTimestamp)
	// 2020/09/04 19:39:26 templated-namespacetemplated-deployment 000099af-14e4-4c1c-928b-581df9e1b0fa 2020-09-04 19:39:27 +0200 CEST

	// perform get on Deployment deleting it - because of registered hook
	_, err = controller.GetKind(Deployment)
	if err != nil {
		panic(err)
	}

	// now make sure its deleted
	_, err = controller.GetKind(Deployment)
	if err == nil {
		panic("its not deleted")
	}

}
