"""
Imitates the commands defined in **makefile**.

Usage:
- `$ make build`: `python -m make build`
- `$ make run`: `python -m make run`
- `$ make clean`: `python -m make clean`
"""

from pathlib import Path


BINARY_NAME = Path("main.out")

SCRIPTS = Path("client/scripts")
STYLES = Path("client/styles")
TEMPLATES = Path("server/templates")
