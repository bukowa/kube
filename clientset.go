package kube

import (
	"k8s.io/client-go/kubernetes"
	typedv1apps "k8s.io/client-go/kubernetes/typed/apps/v1"
	typedv1core "k8s.io/client-go/kubernetes/typed/core/v1"
	typedv1beta1net "k8s.io/client-go/kubernetes/typed/networking/v1beta1"
)

// NewClientSet creates new clientset for a given namespace
func NewClientSet(namespace string, client *kubernetes.Clientset) *clientset {
	return &clientset{
		namespace: namespace,
		client:    client,
	}
}

// clientset implements ClientSet
type clientset struct {
	namespace string
	client    *kubernetes.Clientset
}

func (c *clientset) Namespace() string {
	return c.namespace
}

func (c *clientset) Deployments() typedv1apps.DeploymentInterface {
	return c.client.AppsV1().Deployments(c.namespace)
}

func (c *clientset) Ingresses() typedv1beta1net.IngressInterface {
	return c.client.NetworkingV1beta1().Ingresses(c.namespace)
}

func (c *clientset) Secrets() typedv1core.SecretInterface {
	return c.client.CoreV1().Secrets(c.namespace)
}

func (c *clientset) Services() typedv1core.ServiceInterface {
	return c.client.CoreV1().Services(c.namespace)
}

func (c *clientset) ConfigMaps() typedv1core.ConfigMapInterface {
	return c.client.CoreV1().ConfigMaps(c.namespace)
}

func (c *clientset) PersistentVolumeClaims() typedv1core.PersistentVolumeClaimInterface {
	return c.client.CoreV1().PersistentVolumeClaims(c.namespace)
}
