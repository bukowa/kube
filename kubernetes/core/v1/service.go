package v1

import (
	"context"
	"github.com/bukowa/kube"
	v1core "k8s.io/api/core/v1"
	v1meta "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Service string

func (k Service) Delete(contr kube.ClientSet, ctx context.Context, name string, opts v1meta.DeleteOptions) error {
	return contr.Services().Delete(ctx, name, opts)
}

func (k Service) Get(contr kube.ClientSet, ctx context.Context, name string, opts v1meta.GetOptions) (kube.Resource, error) {
	return contr.Services().Get(ctx, name, opts)
}

func (k Service) Create(contr kube.ClientSet, ctx context.Context, res kube.Resource, opts v1meta.CreateOptions) (kube.Resource, error) {
	if v, ok := res.(*v1core.Service); ok {
		return contr.Services().Create(ctx, v, opts)
	}
	return nil, kube.ErrorInvalidTypeCreate(k)
}

func (k Service) Name() string        { return string(k) }
func (k Service) Cast() kube.Resource { return &v1core.Service{} }
