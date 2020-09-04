package kube

import "fmt"

type ErrorInvalidTypeCreate string

func (e ErrorInvalidTypeCreate) Error() string {
	return fmt.Sprintf("invalid type create for %s", string(e))
}
