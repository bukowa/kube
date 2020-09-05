package v1

import (
	"context"
	"github.com/bukowa/kube"
	"github.com/bukowa/kube/kubernetes"
	v1core "k8s.io/api/core/v1"
	v1meta "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Secret string

func (k Secret) Name() string        { return string(k) }
func (k Secret) Cast() kube.Resource { return &v1core.Secret{} }

func (k Secret) Delete(client kube.ClientSet, ctx context.Context, name string, opts v1meta.DeleteOptions) error {
	return client.Secrets().Delete(ctx, name, opts)
}

func (k Secret) Get(client kube.ClientSet, ctx context.Context, name string, opts v1meta.GetOptions) (kube.Resource, error) {
	return client.Secrets().Get(ctx, name, opts)
}

func (k Secret) Create(client kube.ClientSet, ctx context.Context, res kube.Resource, opts v1meta.CreateOptions) (kube.Resource, error) {
	if v, ok := res.(*v1core.Secret); ok {
		return client.Secrets().Create(ctx, v, opts)
	}
	return nil, kubernetes.ErrorInvalidTypeCreate(k)
}
