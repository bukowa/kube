package v1

import (
	"context"
	"github.com/bukowa/kube"
	"github.com/bukowa/kube/kubernetes"
	v1core "k8s.io/api/core/v1"
	v1meta "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type PersistentVolumeClaim string

func (k PersistentVolumeClaim) Delete(contr kube.BasicClientSet, ctx context.Context, name string, opts v1meta.DeleteOptions) error {
	return contr.PersistentVolumeClaims().Delete(ctx, name, opts)
}

func (k PersistentVolumeClaim) Get(contr kube.BasicClientSet, ctx context.Context, name string, opts v1meta.GetOptions) (kube.Resource, error) {
	return contr.PersistentVolumeClaims().Get(ctx, name, opts)
}

func (k PersistentVolumeClaim) Create(contr kube.BasicClientSet, ctx context.Context, res kube.Resource, opts v1meta.CreateOptions) (kube.Resource, error) {
	if v, ok := res.(*v1core.PersistentVolumeClaim); ok {
		return contr.PersistentVolumeClaims().Create(ctx, v, opts)
	}
	return nil, kubernetes.ErrorInvalidTypeCreate(k)
}

func (k PersistentVolumeClaim) Cast() kube.Resource { return &v1core.PersistentVolumeClaim{} }
func (k PersistentVolumeClaim) Name() string        { return string(k) }
