package kube

type HookType string
type HookErrType string

const (
	PreCreate  HookType = "pre_create"
	PostCreate HookType = "post_create"

	PreGet  HookType = "pre_get"
	PostGet HookType = "post_get"

	PreDelete  HookType = "pre_delete"
	PostDelete HookType = "post_delete"

	ErrCreate HookErrType = "err_create"
	ErrGet    HookErrType = "err_get"
	ErrDelete HookErrType = "err_delete"
)

// Hook is a hook for a Container
// for ex. hooks.SetLabels sets labels for all Resource's in a Container
type Hook func(Container) error

// HookMap is a map of slice of Hook executed for specified Kind
type Hooks map[HookType][]Hook

// HookErr is a hook that happens when an error occurs during get/create/delete
type HookErr func(Container, Kind, error) error

// HooksErr is a map of HookErr executed when an error occurs
type HooksErr map[HookErrType][]HookErr
