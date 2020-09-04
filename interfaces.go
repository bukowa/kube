package kube

import (
	"context"
	"fmt"
	v1apps "k8s.io/api/apps/v1"
	v1core "k8s.io/api/core/v1"
	v1beta1net "k8s.io/api/networking/v1beta1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/kubernetes"

	v1meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	typedv1apps "k8s.io/client-go/kubernetes/typed/apps/v1"
	typedv1core "k8s.io/client-go/kubernetes/typed/core/v1"
	typedv1beta1net "k8s.io/client-go/kubernetes/typed/networking/v1beta1"
)

// Resource represents kubernetes resource
// grouping common interfaces that they implement
type Resource interface {
	v1meta.Common
	v1meta.Object
	v1meta.ObjectMetaAccessor
	schema.ObjectKind
	fmt.Stringer
}

// Kind is meant to be declared as a constant, that is - a string
// that allows accessing the underlying kubernetes state of the
// Resource in a typed (constant) way via Container.
type Kind interface {
	Name() string
	Cast() Resource // Cast() should return new instance of Resource
	Delete(contr ClientSet, ctx context.Context, name string, opts v1meta.DeleteOptions) error
	Get(contr ClientSet, ctx context.Context, name string, opts v1meta.GetOptions) (Resource, error)
	Create(contr ClientSet, ctx context.Context, res Resource, opts v1meta.CreateOptions) (Resource, error)
}

// Container is meant to hold binding between Kind and it's Resource
type Container interface {
	Self() Container                // Self() returns the Container instance itself
	Copy() Container                // Copy() should return a copy of Container so it can be declared once per package
	Update(Kind, Resource) error    // Update() overrides underlying Kind Resource
	GetResource(Kind) Resource      // GetResource() returns underlying Resource for a Kind
	ForEachResource(func(Resource)) // ForEachResource() performs function on each Resource
	ForEachKind(func(Kind))         // ForEachKind() performs function on each Kind

	// caster methods to get underlying Resource instance for a Kind
	Namespace(Kind) *v1core.Namespace
	Deployment(Kind) *v1apps.Deployment
	Ingress(Kind) *v1beta1net.Ingress
	Secret(Kind) *v1core.Secret
	Service(Kind) *v1core.Service
	ConfigMap(Kind) *v1core.ConfigMap
	PersistentVolumeClaim(Kind) *v1core.PersistentVolumeClaim
}

// ResourcesManager is meant to manipulate Resource state in kubernetes
type ClientSet interface {
	Namespace() string
	Client() *kubernetes.Clientset
	Namespaces() typedv1core.NamespaceInterface
	Deployments() typedv1apps.DeploymentInterface
	Ingresses() typedv1beta1net.IngressInterface
	Secrets() typedv1core.SecretInterface
	Services() typedv1core.ServiceInterface
	ConfigMaps() typedv1core.ConfigMapInterface
	PersistentVolumeClaims() typedv1core.PersistentVolumeClaimInterface
}

// Controller is meant to perform actions on Container via ClientSet
// a Controller can use Container New() method to construct itself - that allows declaring Container once per package.
// Each method on Controller handles the underlying Container - updating Resource's that it holds.
type Controller interface {
	Container // Container of which Controller is in control

	RegisterHooks(Hooks) // RegisterHooks accepts hooks that Controller may use
	ClientSet() ClientSet

	// todo should hooks run here?
	GetKind(Kind) (Resource, error)    // GetKind(Kind) performs kubernetes get request updating underlying Resource
	CreateKind(Kind) (Resource, error) // CreateKind(Kind) performs kubernetes create request for underlying Resource
	DeleteKind(Kind) error             // DeleteKind(Kind) performs kubernetes delete request for underlying Resource

	GetContainer() []error    // GetContainer() performs kubernetes get request for all Kind's in Container
	CreateContainer() []error // CreateContainer() performs kubernetes create request for all Kind's in Container
	DeleteContainer() []error // DeleteContainer() performs kubernetes delete request for all Kind's in Container
}
