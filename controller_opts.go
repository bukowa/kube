package kube

import v1meta "k8s.io/apimachinery/pkg/apis/meta/v1"

type Option func(*controller)

var WithClientSet = func(clientset ClientSet) Option {
	return func(c *controller) {
		c.clientset = clientset
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
