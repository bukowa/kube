package kube

// Hook is a generic hook for all Resource types
// for example look at hooks.SetLabels that sets labels
// for all Resource's in a Container
type Hook func(Container) error

// HookKind is a hook for specified Kind
type HookKind func(ClientSet, Kind, Resource) error

type HookKindMap map[Kind][]HookKind
