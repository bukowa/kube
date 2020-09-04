package v1

import (
	"context"
	"github.com/bukowa/kube"
	"github.com/bukowa/kube/kubernetes"
	v1core "k8s.io/api/core/v1"
	v1meta "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ConfigMap string

func (k ConfigMap) Delete(contr kube.ClientSet, ctx context.Context, name string, opts v1meta.DeleteOptions) error {
	return contr.ConfigMaps().Delete(ctx, name, opts)
}

func (k ConfigMap) Get(contr kube.ClientSet, ctx context.Context, name string, opts v1meta.GetOptions) (kube.Resource, error) {
	return contr.ConfigMaps().Get(ctx, name, opts)
}

func (k ConfigMap) Create(contr kube.ClientSet, ctx context.Context, res kube.Resource, opts v1meta.CreateOptions) (kube.Resource, error) {
	if v, ok := res.(*v1core.ConfigMap); ok {
		return contr.ConfigMaps().Create(ctx, v, opts)
	}
	return nil, kubernetes.ErrorInvalidTypeCreate(k)
}

func (k ConfigMap) Name() string        { return string(k) }
func (k ConfigMap) Cast() kube.Resource { return &v1core.ConfigMap{} }
