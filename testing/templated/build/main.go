// check how templates behave when application is compiled

package main

import (
	"fmt"
	"github.com/bukowa/kube/templated"
	test "github.com/bukowa/kube/testing/templated"
	"log"
)

func main() {
	c := test.Container
	d := test.Data

	if err := templated.TemplateHook(c, d)(c); err != nil {
		panic(err)
	}
	// for each kind check its name
	for _, kind := range test.Group {
		if resource := c.GetResource(kind); resource == nil {
			panic(fmt.Sprintf("resource: %s is nil", kind.Name()))
		} else if resource.GetName() != "new" {
			panic(fmt.Sprintf("resource: %s name is wrong", kind.Name()))
		}
	}
	log.Print("success")
}
