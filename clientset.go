package kube

import (
	"k8s.io/client-go/kubernetes"
	typedv1apps "k8s.io/client-go/kubernetes/typed/apps/v1"
	typedv1core "k8s.io/client-go/kubernetes/typed/core/v1"
	typedv1beta1net "k8s.io/client-go/kubernetes/typed/networking/v1beta1"
)

// NewClientSet creates new KubeClientSet for a given namespace
func NewClientSet(namespace string, client *kubernetes.Clientset) ClientSet {
	return &BasicClientSet{
		namespace: namespace,
		client:    client,
	}
}

// KubeClientSet implements BasicClientSet
type BasicClientSet struct {
	namespace string
	client    *kubernetes.Clientset
}

func (c *BasicClientSet) Namespace() string {
	return c.namespace
}

func (c *BasicClientSet) Client() *kubernetes.Clientset {
	return c.client
}

func (c *BasicClientSet) Namespaces() typedv1core.NamespaceInterface {
	return c.client.CoreV1().Namespaces()
}

func (c *BasicClientSet) Deployments() typedv1apps.DeploymentInterface {
	return c.client.AppsV1().Deployments(c.namespace)
}

func (c *BasicClientSet) Ingresses() typedv1beta1net.IngressInterface {
	return c.client.NetworkingV1beta1().Ingresses(c.namespace)
}

func (c *BasicClientSet) Secrets() typedv1core.SecretInterface {
	return c.client.CoreV1().Secrets(c.namespace)
}

func (c *BasicClientSet) Services() typedv1core.ServiceInterface {
	return c.client.CoreV1().Services(c.namespace)
}

func (c *BasicClientSet) ConfigMaps() typedv1core.ConfigMapInterface {
	return c.client.CoreV1().ConfigMaps(c.namespace)
}

func (c *BasicClientSet) PersistentVolumeClaims() typedv1core.PersistentVolumeClaimInterface {
	return c.client.CoreV1().PersistentVolumeClaims(c.namespace)
}
