package v1

import (
	"context"
	"github.com/bukowa/kube"
	"github.com/bukowa/kube/kubernetes"
	v1core "k8s.io/api/core/v1"
	v1meta "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Secret string

func (k Secret) Delete(contr kube.BasicClientSet, ctx context.Context, name string, opts v1meta.DeleteOptions) error {
	return contr.Secrets().Delete(ctx, name, opts)
}

func (k Secret) Get(contr kube.BasicClientSet, ctx context.Context, name string, opts v1meta.GetOptions) (kube.Resource, error) {
	return contr.Secrets().Get(ctx, name, opts)
}

func (k Secret) Create(contr kube.BasicClientSet, ctx context.Context, res kube.Resource, opts v1meta.CreateOptions) (kube.Resource, error) {
	if v, ok := res.(*v1core.Secret); ok {
		return contr.Secrets().Create(ctx, v, opts)
	}
	return nil, kubernetes.ErrorInvalidTypeCreate(k)
}

func (k Secret) Name() string        { return string(k) }
func (k Secret) Cast() kube.Resource { return &v1core.Secret{} }
