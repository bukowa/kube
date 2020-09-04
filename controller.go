package kube

import (
	"context"
	v1meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type Option func(*controller)

// NewController creates new controller
func NewController(namespace string, container Container, client *kubernetes.Clientset, opts ...Option) *controller {
	c := &controller{
		clientset:  NewClientSet(namespace, client),
		container:  container,
		CreateOpts: v1meta.CreateOptions{},
		DeleteOpts: v1meta.DeleteOptions{},
		GetOpts:    v1meta.GetOptions{},
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

// controller implements Controller
type controller struct {
	clientset  ClientSet
	container  Container
	CreateOpts v1meta.CreateOptions
	DeleteOpts v1meta.DeleteOptions
	GetOpts    v1meta.GetOptions
}

func (c controller) Container() Container {
	return c.container
}

func (c *controller) Create() []error {
	return c.create(context.Background())
}

func (c *controller) Delete() []error {
	return c.delete(context.Background())
}

func (c *controller) Get() []error {
	return c.get(context.Background())
}

func (c *controller) create(ctx context.Context) (errs []error) {
	var handle = func(k Kind, r Resource, e error) {
		if e != nil {
			errs = append(errs, e)
			return
		}
		if e = c.container.Update(k, r); e != nil {
			errs = append(errs, e)
		}
	}
	c.container.ForEachKind(func(kind Kind) {
		res, err := kind.Create(c.clientset, ctx, c.container.GetResource(kind), c.CreateOpts)
		handle(kind, res, err)
	})
	return
}

func (c *controller) delete(ctx context.Context) (errs []error) {
	var handle = func(e error) {
		if e != nil {
			errs = append(errs, e)
		}
	}
	c.container.ForEachKind(func(kind Kind) {
		handle(kind.Delete(c.clientset, ctx, c.container.GetResource(kind).GetName(), c.DeleteOpts))
	})
	return
}

func (c *controller) get(ctx context.Context) (errs []error) {
	var handle = func(k Kind, r Resource, e error) {
		if e != nil {
			errs = append(errs, e)
			return
		}
		if e = c.container.Update(k, r); e != nil {
			errs = append(errs, e)
		}
	}
	c.container.ForEachKind(func(kind Kind) {
		res, err := kind.Get(c.clientset, ctx, c.container.GetResource(kind).GetName(), c.GetOpts)
		handle(kind, res, err)
	})
	return
}
