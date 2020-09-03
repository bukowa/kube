import utils


def test_get_imports():
    with open("testdata/deployment.go") as f:
        assert utils.get_imports(f) == [
            ("", "context"),
            ("", "time"),
            ("v1", "k8s.io/api/apps/v1"),
            ("autoscalingv1", "k8s.io/api/autoscaling/v1"),
            ("metav1", "k8s.io/apimachinery/pkg/apis/meta/v1"),
            ("types", "k8s.io/apimachinery/pkg/types"),
            ("watch", "k8s.io/apimachinery/pkg/watch"),
            ("scheme", "k8s.io/client-go/kubernetes/scheme"),
            ("rest", "k8s.io/client-go/rest"),
        ]