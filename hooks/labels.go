package hooks

import "github.com/bukowa/kube"

// HookSetLabels overrides Labels
var SetLabels = func(labels map[string]string) kube.Hook {
	return func(manager kube.Container) error {
		return manager.ForEachResource(func(res kube.Resource) error {
			res.SetLabels(labels)
			return nil
		})
	}
}

// HookUpdateLabels updates Labels with new values
var UpdateLabels = func(labels map[string]string) kube.Hook {
	return func(manager kube.Container) error {
		return manager.ForEachResource(func(res kube.Resource) error {
			res.SetLabels(updateMapString(res.GetLabels(), labels))
			return nil
		})
	}
}

type labels map[string]string

func updateMapString(existing, new labels) labels {
	for k, v := range new {
		existing[k] = v
	}
	return existing
}
