package v1

import (
	"context"
	"github.com/bukowa/kube"
	"github.com/bukowa/kube/kubernetes"
	v1core "k8s.io/api/core/v1"
	v1meta "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Namespace string

func (k Namespace) Name() string {
	return string(k)
}

func (k Namespace) Cast() kube.Resource {
	return &v1core.Namespace{}
}

func (k Namespace) Delete(contr kube.BasicClientSet, ctx context.Context, name string, opts v1meta.DeleteOptions) error {
	return contr.Namespaces().Delete(ctx, name, opts)
}

func (k Namespace) Get(contr kube.BasicClientSet, ctx context.Context, name string, opts v1meta.GetOptions) (kube.Resource, error) {
	return contr.Namespaces().Get(ctx, name, opts)
}

func (k Namespace) Create(contr kube.BasicClientSet, ctx context.Context, res kube.Resource, opts v1meta.CreateOptions) (kube.Resource, error) {
	if v, ok := res.(*v1core.Namespace); ok {
		return contr.Namespaces().Create(ctx, v, opts)
	}
	return nil, kubernetes.ErrorInvalidTypeCreate(k)
}
