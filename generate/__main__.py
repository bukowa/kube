import os
from dataclasses import dataclass
from typing import List


@dataclass
class Param:
    name: str       # ctx
    type: str       # context.Context


@dataclass
class Method:
    name: str               # Create
    params: List[Param]     # (ctx context.Context, deployment *v1beta1.Deployment, opts v1.CreateOptions)
    results: List[Param]    # (*v1beta1.Deployment, error)


@dataclass
class Interface:
    name: str               # DeploymentInterface
    methods: List[Method]   # Create Update UpdateStatus Delete DeleteCollection


@dataclass
class Typed:
    group: str      # apps
    version: str    # v1beta1


if __name__ == '__main__':
    for root, dirs, files in os.walk("../vendor/k8s.io/client-go/kubernetes/typed/"):
        print(root, dirs, files)
