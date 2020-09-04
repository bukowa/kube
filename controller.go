package kube

import (
	"context"
	v1meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// NewController creates new Controller
func NewController(container Container, hooks Hooks, opts ...Option) Controller {
	if hooks == nil {
		hooks = make(Hooks)
	}
	c := &controller{
		Container:  container,
		CreateOpts: v1meta.CreateOptions{},
		DeleteOpts: v1meta.DeleteOptions{},
		GetOpts:    v1meta.GetOptions{},
		hooks:      hooks,
	}
	for _, opt := range opts {
		opt(c)
	}
	if c.clientset == nil {
		panic("clientset is nil, you can provide some with opts")
	}
	return c
}

type Option func(*controller)

var WithKubernetesClient = func(namespace string, client *kubernetes.Clientset) Option {
	return func(c *controller) {
		c.clientset = NewClientSet(namespace, client)
	}
}

var WithClientSet = func(clientset ClientSet) Option {
	return func(c *controller) {
		c.clientset = clientset
	}
}

var CreateOpts = func(opts v1meta.CreateOptions) Option {
	return func(c *controller) {
		c.CreateOpts = opts
	}
}

// OptionContainerCopy creates controller copying the Container
var ContainerCopy = func() Option {
	return func(c *controller) {
		c.Container = c.Container.Copy()
	}
}

// controller implements Controller
type controller struct {
	Container

	clientset  ClientSet
	CreateOpts v1meta.CreateOptions
	DeleteOpts v1meta.DeleteOptions
	GetOpts    v1meta.GetOptions
	hooks      Hooks
}

// RegisterHooks
func (c *controller) RegisterHooks(hooks Hooks) {
	for t, h := range hooks {
		c.hooks[t] = append(c.hooks[t], h...)
	}
}

func (c *controller) Hooks() Hooks {
	return c.hooks
}

func (c *controller) GetKind(kind Kind) (Resource, error) {
	return kind.Get(c.clientset, context.Background(), c.GetResource(kind).GetName(), c.GetOpts)
}

func (c *controller) DeleteKind(kind Kind) error {
	return kind.Delete(c.clientset, context.Background(), c.GetResource(kind).GetName(), c.DeleteOpts)
}

func (c *controller) ClientSet() ClientSet {
	return c.clientset
}

func (c *controller) CreateKind(kind Kind) (res Resource, err error) {
	if err = runHooks(c, PreCreate); err != nil {
		return
	}
	if res, err = kind.Create(c.clientset, context.Background(), c.GetResource(kind), c.CreateOpts); err != nil {
		return
	}
	return res, runHooks(c, PostCreate)
}

func (c *controller) CreateContainer() []error {
	return c.create(context.Background())
}

func (c *controller) DeleteContainer() []error {
	return c.delete(context.Background())
}

func (c *controller) GetContainer() []error {
	return c.get(context.Background())
}

func (c *controller) create(ctx context.Context) (errs []error) {
	// PreCreate hook
	if err := runHooks(c, PreCreate); err != nil {
		return []error{err}
	}
	// helper function
	var handle = func(k Kind, r Resource, e error) {
		if e != nil {
			errs = append(errs, e)
			return
		}
		if e = c.Update(k, r); e != nil {
			errs = append(errs, e)
		}
	}
	// create each kind
	c.ForEachKind(func(kind Kind) {
		res, err := kind.Create(c.clientset, ctx, c.GetResource(kind), c.CreateOpts)
		handle(kind, res, err)
	})
	// PostCreate hook
	if err := runHooks(c, PostCreate); err != nil {
		errs = append(errs, err)
	}
	return errsReturn(errs)
}

func (c *controller) delete(ctx context.Context) (errs []error) {

	var handle = func(e error) {
		if e != nil {
			errs = append(errs, e)
		}
	}

	// pre hook
	if err := runHooks(c, PreDelete); err != nil {
		return []error{err}
	}
	// delete all
	c.ForEachKind(func(kind Kind) {
		handle(kind.Delete(c.clientset, ctx, c.GetResource(kind).GetName(), c.DeleteOpts))
	})
	// post hook
	if err := runHooks(c, PostDelete); err != nil {
		errs = append(errs, err)
	}
	return errsReturn(errs)
}

func (c *controller) get(ctx context.Context) (errs []error) {
	var handle = func(k Kind, r Resource, e error) {
		if e != nil {
			errs = append(errs, e)
			return
		}
		if e = c.Update(k, r); e != nil {
			errs = append(errs, e)
		}
	}
	// pre hook
	if err := runHooks(c, PreGet); err != nil {
		return []error{err}
	}

	// get all
	c.ForEachKind(func(kind Kind) {
		res, err := kind.Get(c.clientset, ctx, c.GetResource(kind).GetName(), c.GetOpts)
		handle(kind, res, err)
	})

	// post hook
	if err := runHooks(c, PostGet); err != nil {
		errs = append(errs, err)
	}
	return errsReturn(errs)
}

func errsReturn(errs []error) []error {
	if len(errs) == 0 {
		return nil
	}
	return errs
}

func runHooks(c *controller, t HookType) error {
	if c.hooks[t] == nil {
		return nil
	}
	for _, hook := range c.hooks[t] {
		if err := hook(c); err != nil {
			return err
		}
	}
	return nil
}
