package v1beta1

import (
	"context"
	"github.com/bukowa/kube"
	"github.com/bukowa/kube/kubernetes"
	v1beta1net "k8s.io/api/networking/v1beta1"
	v1meta "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Ingress string

func (k Ingress) Name() string        { return string(k) }
func (k Ingress) Cast() kube.Resource { return &v1beta1net.Ingress{} }

func (k Ingress) Delete(client kube.ClientSet, ctx context.Context, name string, opts v1meta.DeleteOptions) error {
	return client.Ingresses().Delete(ctx, name, opts)
}

func (k Ingress) Get(client kube.ClientSet, ctx context.Context, name string, opts v1meta.GetOptions) (kube.Resource, error) {
	return client.Ingresses().Get(ctx, name, opts)
}

func (k Ingress) Create(client kube.ClientSet, ctx context.Context, res kube.Resource, opts v1meta.CreateOptions) (kube.Resource, error) {
	if v, ok := res.(*v1beta1net.Ingress); ok {
		return client.Ingresses().Create(ctx, v, opts)
	}
	return nil, kubernetes.ErrorInvalidTypeCreate(k)
}
