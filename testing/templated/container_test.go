package templated

import (
	"github.com/bukowa/kube/templated"
	"testing"
)

func TestContainer(t *testing.T) {
	c := Container
	d := Data

	//
	if d.Name != "new" {
		t.Error()
	}

	// execute templates
	if err := templated.TemplateHook(c, d)(c); err != nil {
		t.Error(err)
	}

	// for each kind check its name
	for _, kind := range Group {
		if resource := c.GetResource(kind); resource == nil {
			t.Errorf("resource: %s is nil", kind.Name())
		} else if resource.GetName() != "new" {
			t.Errorf("resource: %s name is wrong", kind.Name())
		}
	}
}
