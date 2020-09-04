package kube

import (
	v1meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type Option func(*BasicController)

var WithHooks = func(hooks Hooks) Option {
	return func(c *BasicController) {
		for t, h := range hooks {
			if existing := c.hooks[t]; existing == nil {
				c.hooks[t] = h
			} else {
				c.hooks[t] = append(c.hooks[t], h...)
			}
		}
	}
}

var WithClientSet = func(clientset ClientSet) Option {
	return func(c *BasicController) {
		c.clientset = clientset
	}
}

var WithKubernetesClient = func(namespace string, client *kubernetes.Clientset) Option {
	return func(c *BasicController) {
		c.clientset = NewClientSet(namespace, client)
	}
}

var WithCreateOpts = func(opts v1meta.CreateOptions) Option {
	return func(c *BasicController) {
		c.CreateOpts = opts
	}
}

// OptionContainerCopy creates BasicController copying the Container
var WithContainerCopy = func() Option {
	return func(c *BasicController) {
		c.Container = c.Container.Copy()
	}
}
