package kube

import (
	"context"
	v1core "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"path"
)

var namespace = "asd9asda8sd9as9dxzasdjiasdsa8d7sa87as87d8asjasdskja"
var Namespace = &v1core.Namespace{
	ObjectMeta: v1.ObjectMeta{Name: namespace},
}

func checkNamespace(c *kubernetes.Clientset) {
	if _, err := c.CoreV1().Namespaces().Get(context.TODO(), namespace, v1.GetOptions{}); err == nil {
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
