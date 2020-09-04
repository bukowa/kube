package templated

import (
	"github.com/bukowa/kube"
)

type Controller struct {
	kube.Controller
	container *Container
}

// NewController creates Controller
func NewController(container *Container, data interface{}, hooks kube.Hooks, opts ...kube.Option) *Controller {
	c := &Controller{
		Controller: kube.NewController(container, nil, opts...),
	}
	c.container = c.Self().(*Container)
	// register hooks in order
	c.RegisterHooks(Hooks(c.container, data))
	c.RegisterHooks(hooks)
	return c
}
