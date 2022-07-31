"""
Imitates the `$ make build` command.
"""

import os

from .constants import BINARY_NAME, SCRIPTS, STYLES, TEMPLATES


def main():
    """
    Runs the commands defined in **makefile** for the `$ make build` command.
    """
    # Compile/Transpile TypeScript source files
    for file in SCRIPTS.glob("*.ts"):
        os.system(f"tsc {file}")

    # Compile/Transpile SASS/SCSS source files
    os.system(f"sass {STYLES}/:{STYLES}/")
    os.system(f"sass {TEMPLATES}/:{TEMPLATES}/")

    # Compile Go package
    os.system(f"go build -o {BINARY_NAME} main.go")


if __name__ == "__main__":
    main()
