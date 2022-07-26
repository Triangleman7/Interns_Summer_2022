"""
Imitates the `$ make run` command.
"""

import os
import subprocess
import sys

from . import BINARY_NAME, SCRIPTS, STYLES, TEMPLATES


def main():
    for file in SCRIPTS.glob("*.ts"):
        os.system(f"tsc {file}")
    os.system(f"sass {STYLES}/:{STYLES}/")
    os.system(f"sass {TEMPLATES}/:{TEMPLATES}/")
    os.system(f"go build -o {BINARY_NAME} main.go")
    try:
        os.system(f"{BINARY_NAME}")
    except KeyboardInterrupt:
        sys.exit(0)
