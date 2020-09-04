package kube

import (
	"context"
	"fmt"
	v1apps "k8s.io/api/apps/v1"
	v1core "k8s.io/api/core/v1"
	v1beta1net "k8s.io/api/networking/v1beta1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	typedv1apps "k8s.io/client-go/kubernetes/typed/apps/v1"
	typedv1core "k8s.io/client-go/kubernetes/typed/core/v1"
	typedv1beta1net "k8s.io/client-go/kubernetes/typed/networking/v1beta1"
)

// Resource represents kubernetes resource -
// - grouping common interfaces that they implement
type Resource interface {
	v1meta.Common
	v1meta.Object
	v1meta.ObjectMetaAccessor
	schema.ObjectKind
	fmt.Stringer
}

// Kind is meant to be declared as a constant, that is - a string
// that allows accessing the underlying kubernetes state of the
// Resource in a typed (constant) way via ResourcesConverter.
type Kind interface {
	Name() string
	Cast() Resource
	Delete(contr ClientSet, ctx context.Context, name string, opts v1meta.DeleteOptions) error
	Get(contr ClientSet, ctx context.Context, name string, opts v1meta.GetOptions) (Resource, error)
	Create(contr ClientSet, ctx context.Context, res Resource, opts v1meta.CreateOptions) (Resource, error)
}

// Container is meant to hold a binding between Kind and it's Resource
type Container interface {
	Update(Kind, Resource) error
	GetResource(Kind) Resource
	ForEachResource(func(Resource))
	ForEachKind(func(Kind))
	Deployment(Kind) *v1apps.Deployment
	Ingress(Kind) *v1beta1net.Ingress
	Secret(Kind) *v1core.Secret
	Service(Kind) *v1core.Service
	ConfigMap(Kind) *v1core.ConfigMap
	PersistentVolumeClaim(Kind) *v1core.PersistentVolumeClaim
}

// ResourcesManager is meant to
// ClientSet is used to manipulate Resource state in kubernetes
type ClientSet interface {
	Deployments() typedv1apps.DeploymentInterface
	Ingresses() typedv1beta1net.IngressInterface
	Secrets() typedv1core.SecretInterface
	Services() typedv1core.ServiceInterface
	ConfigMaps() typedv1core.ConfigMapInterface
	PersistentVolumeClaims() typedv1core.PersistentVolumeClaimInterface
}

type Controller interface {
	Create() []error
	Delete() []error
	Get() []error
}
