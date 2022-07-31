"""
Runs an imitation of the `$ make (build|clean|run|test)` command.

Usage:
- `$ python -m make (build|clean|run|test)`
"""

import sys

from . import build, clean, run, test


COMMANDS = {
    "build": build.main,
    "clean": clean.main,
    "run": run.main,
    "test": test.main
}


if __name__ == "__main__":
    cmd = COMMANDS.get(sys.argv[1])
    cmd()
