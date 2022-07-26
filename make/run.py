"""
Imitates the `$ make run` command.
"""

import os
import subprocess
import sys

from .constants import BINARY_NAME, SCRIPTS, STYLES, TEMPLATES


def main():
    # Compile/Transpile TypeScript source files
    for file in SCRIPTS.glob("*.ts"):
        os.system(f"tsc {file}")
    # Compile/Transpile SASS/SCSS source files
    os.system(f"sass {STYLES}/:{STYLES}/")
    os.system(f"sass {TEMPLATES}/:{TEMPLATES}/")

    # Compile Go package
    os.system(f"go build -o {BINARY_NAME} main.go")
    try:
        os.system(f"{BINARY_NAME}")
    except KeyboardInterrupt:
        sys.exit(0)
