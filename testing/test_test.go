package testing

import (
	"context"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"path"
)

func checkNamespace(n string, c *kubernetes.Clientset) {
	if _, err := c.CoreV1().Namespaces().Get(context.TODO(), n, v1.GetOptions{}); err == nil {
		panic("namespace already exists")
	}
}

func HomeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE")
}

func GetKubernetesClient() (*kubernetes.Clientset, error) {
	config, err := clientcmd.BuildConfigFromFlags("", path.Join(HomeDir(), ".kube", "config"))
	if err != nil {
		return nil, err
	}
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	return client, nil
}
