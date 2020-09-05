package templated

import (
	"github.com/bukowa/kube"
)

type Controller struct {
	kube.Controller
	container *Container
}

// NewController creates Controller
func NewController(container *Container, data interface{}, opts ...kube.Option) *Controller {
	c := &Controller{
		Controller: kube.NewController(container, opts...),
	}
	c.container = c.Self().(*Container)
	// configure hooks to run first
	c.Configure(kube.WithHooksFirst(TemplateHooks(c.container, data)))
	return c
}
