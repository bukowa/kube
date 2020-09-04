package templated

import (
	"fmt"
	"github.com/bukowa/kube"
)

type Controller struct {
	kube.Controller
	container *Container
}

// NewController creates Controller
func NewController(container *Container, data interface{}, hooks kube.Hooks, opts ...kube.Option) *Controller {
	tc := &Controller{
		Controller: kube.NewController(container, hooks, opts...),
	}
	if v, ok := tc.Self().(*Container); !ok {
		tc.container = v
	} else {
		// todo
		panic(fmt.Sprintf("container holds wrong type: %T", v))
	}
	// register basic hooks
	tc.RegisterHooks(Hooks(tc.container, data))
	return tc
}
