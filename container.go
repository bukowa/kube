package kube

import (
	"fmt"
	v1apps "k8s.io/api/apps/v1"
	v1core "k8s.io/api/core/v1"
	v1beta1net "k8s.io/api/networking/v1beta1"
)

// NewContainer creates new container for Kind's
func NewContainer(kinds ...Kind) *container {
	binding := make(map[Kind]Resource, len(kinds))
	for _, kind := range kinds {
		if binding[kind] != nil {
			panic(fmt.Sprintf("duplicate kind %s", kind))
		}
		binding[kind] = kind.Cast()
	}
	return &container{
		kinds:   kinds,
		binding: binding,
	}
}

// container implements Container interface
type container struct {
	kinds   []Kind
	binding map[Kind]Resource
}

func (c *container) Self() Container {
	return c
}

func (c *container) Copy() Container {
	return NewContainer(c.kinds...)
}

func (c *container) Update(kind Kind, resource Resource) error {
	if c.binding[kind] == nil {
		return fmt.Errorf("cannot handle kind %s", kind)
	}
	c.binding[kind] = resource
	return nil
}

func (c *container) GetResource(kind Kind) Resource {
	for k, v := range c.binding {
		if k == kind {
			return v
		}
	}
	return nil
}

func (c *container) ForEachResource(f func(Resource)) {
	for _, res := range c.binding {
		f(res)
	}
}

func (c *container) ForEachKind(f func(Kind)) {
	for _, kind := range c.kinds {
		f(kind)
	}
}

func (c *container) Namespace(kind Kind) *v1core.Namespace {
	if res := c.GetResource(kind); res != nil {
		return res.(*v1core.Namespace)
	}
	return nil
}

func (c *container) Deployment(kind Kind) *v1apps.Deployment {
	if res := c.GetResource(kind); res != nil {
		return res.(*v1apps.Deployment)
	}
	return nil
}

func (c *container) Ingress(kind Kind) *v1beta1net.Ingress {
	if res := c.GetResource(kind); res != nil {
		return res.(*v1beta1net.Ingress)
	}
	return nil
}

func (c *container) Secret(kind Kind) *v1core.Secret {
	if res := c.GetResource(kind); res != nil {
		return res.(*v1core.Secret)
	}
	return nil
}

func (c *container) Service(kind Kind) *v1core.Service {
	if res := c.GetResource(kind); res != nil {
		return res.(*v1core.Service)
	}
	return nil
}

func (c *container) ConfigMap(kind Kind) *v1core.ConfigMap {
	if res := c.GetResource(kind); res != nil {
		return res.(*v1core.ConfigMap)
	}
	return nil
}

func (c *container) PersistentVolumeClaim(kind Kind) *v1core.PersistentVolumeClaim {
	if res := c.GetResource(kind); res != nil {
		return res.(*v1core.PersistentVolumeClaim)
	}
	return nil
}
