import utils
import pytest


@pytest.fixture()
def file():
    return open("./testdata/deployment.go")


def test_get_package_name(file):
    assert utils.get_package_name(file) == "v1"


def test_get_imports(file):
    assert utils.get_imports(file) == [
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