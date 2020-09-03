import os
from dataclasses import dataclass
from typing import List
import pathlib
import io
import utils

pathlib.Path.

@dataclass
class Import:
    name: str  # v1beta1
    path: str  # "k8s.io/api/apps/v1beta1"

    @classmethod
    def parse(cls, s: str) -> 'Import':
        s.index("import")


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
class File:
    name: str                       # deployment.go
    package: str                    # v1beta1
    imports: List[Import]
    interfaces: List[Interface]

    @classmethod
    def from_file(cls, file: io.TextIOWrapper):
        package = utils.get_package_name(file)
        imports = utils.get_imports(file)

@dataclass
class Folder:
    group: str      # apps
    version: str    # v1beta1
    files: List[File]


if __name__ == '__main__':
    i = 0
    for root, dirs, files in os.walk("../vendor/k8s.io/client-go/kubernetes/typed/"):
        i += 1
        if i == 1:
            continue
        if files:
            _ps = os.path.split(root)
            _l = len(_ps)

            group, version, fileobjects = _ps[_l-1], _ps[_l], []
            for f in files:
                with open(os.path.join(root, f)) as ioFile:
                    fileobjects.append(File(f, ))