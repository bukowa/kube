package kube

type HookType string

const (
	PreCreate  HookType = "pre_create"
	PostCreate HookType = "post_create"

	PreGet  HookType = "pre_get"
	PostGet HookType = "post_get"

	PreDelete  HookType = "pre_delete"
	PostDelete HookType = "post_delete"
)

// Hook is a hook for a Container
// for ex. hooks.SetLabels sets labels for all Resource's in a Container
type Hook func(Container) error

// HookMap is a map of Hook's executed for specified Kind
type Hooks map[HookType][]Hook
