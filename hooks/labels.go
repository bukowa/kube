package hooks

import "github.com/bukowa/kube"

// SetAnnotations overrides Annotations
var SetAnnotations = func(annotations map[string]string) kube.Hook {
	return func(c kube.Container) error {
		c.ForEachResource(func(r kube.Resource) {
			r.SetAnnotations(annotations)
		})
		return nil
	}
}

// UpdateAnnotations updates Annotations with new values
var UpdateAnnotations = func(annotations map[string]string) kube.Hook {
	return func(c kube.Container) error {
		c.ForEachResource(func(r kube.Resource) {
			r.SetAnnotations(updateMapString(r.GetAnnotations(), annotations))
		})
		return nil
	}
}

// HookSetLabels overrides Labels
var SetLabels = func(labels map[string]string) kube.Hook {
	return func(cont kube.Container) error {
		cont.ForEachResource(func(res kube.Resource) {
			res.SetLabels(labels)
		})
		return nil
	}
}

// HookUpdateLabels updates Labels with new values
var UpdateLabels = func(labels map[string]string) kube.Hook {
	return func(cont kube.Container) error {
		cont.ForEachResource(func(res kube.Resource) {
			res.SetLabels(updateMapString(res.GetLabels(), labels))
		})
		return nil
	}
}

type labels map[string]string

func updateMapString(existing, new labels) labels {
	if existing == nil {
		existing = make(map[string]string, len(new))
	}
	for k, v := range new {
		existing[k] = v
	}
	return existing
}
