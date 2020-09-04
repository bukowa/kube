package kube

import (
	v1meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type Option func(*controller)

var WithClientSet = func(clientset BasicClientSet) Option {
	return func(c *controller) {
		c.clientset = clientset
	}
}

var WithKubernetesClient = func(namespace string, client *kubernetes.Clientset) Option {
	return func(c *controller) {
		c.clientset = NewClientSet(namespace, client)
	}
}

var CreateOpts = func(opts v1meta.CreateOptions) Option {
	return func(c *controller) {
		c.CreateOpts = opts
	}
}

// OptionContainerCopy creates controller copying the Container
var ContainerCopy = func() Option {
	return func(c *controller) {
		c.Container = c.Container.Copy()
	}
}
