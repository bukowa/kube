import io
from typing import List, Tuple


def get_package_name(f: io.TextIOWrapper) -> str:
    for line in f:
        if line.startswith("package"):
            return line[len("package")+1:].strip()
    raise Exception(f)


def get_imports(f: io.TextIOWrapper) -> List[Tuple[str, str]]:
    data = [x for x in f]
    for i, line in enumerate(data[:]):
        if line.startswith("import ("):
            data = data[i+1:]
            for _i, _line in enumerate(data[:]):
                if _line.strip() == ")":
                    data = data[:_i]
                    break

    ret = []
    for i, x in enumerate(data[:]):
        splitted = x.strip().split(" ")
        if len(splitted) == 2:
            ret.append((splitted[0], splitted[1].replace('"', "")))
        elif len(splitted) == 1:
            ret.append(("", splitted[0].replace('"', "")))
        else:
            raise Exception(splitted)

    # clear empty
    for i, x in enumerate(ret[:]):
        if x[0] == "" and x[1] == "":
            ret.pop(i)
    return ret


def get_interfaces(f: io.TextIOWrapper):
    data = [x for x in f]
    targets = []
    for i, l in enumerate(data[:]):
        if l.startswith("type"):
            for ii, ll in data[i:]:
                if ll.strip() == "}":
                    # todo
            return