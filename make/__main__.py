"""
Runs an imitation of the `$ make (build|clean|run)` command.

Usage:
- `$ python -m make (build|clean|run)`
"""

import sys

from . import build, clean, run


COMMANDS = {
    "build": build.main,
    "clean": clean.main,
    "run": run.main
}


if __name__ == "__main__":
    cmd = COMMANDS.get(sys.argv[1])
    cmd()
