package v1

import (
	"context"
	"github.com/bukowa/kube"
	v1 "k8s.io/api/apps/v1"
	v1meta "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Deployment string

func (d Deployment) Name() string        { return string(d) }
func (d Deployment) Cast() kube.Resource { return &v1.Deployment{} }

func (d Deployment) Delete(client kube.ClientSet, ctx context.Context, name string, opts v1meta.DeleteOptions) error {
	return client.Deployments().Delete(ctx, name, opts)
}

func (d Deployment) Get(contr kube.ClientSet, ctx context.Context, name string, opts v1meta.GetOptions) (kube.Resource, error) {
	return contr.Deployments().Get(ctx, name, opts)
}

func (d Deployment) Create(contr kube.ClientSet, ctx context.Context, res kube.Resource, opts v1meta.CreateOptions) (kube.Resource, error) {
	if v, ok := res.(*v1.Deployment); ok {
		return contr.Deployments().Create(ctx, v, opts)
	}
	return nil, kube.ErrorInvalidTypeCreate(d)
}
