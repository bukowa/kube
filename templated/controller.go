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
	// register hooks in order
	c.RegisterHooks(TemplateHooks(c.container, data))
	return c
}
