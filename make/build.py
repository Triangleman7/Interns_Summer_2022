"""
Imitates the `$ make build` command.
"""

import os

from .constants import BINARY_NAME, SCRIPTS, STYLES, TEMPLATES


def main():
    for file in SCRIPTS.glob("*.ts"):
        os.system(f"tsc {file}")
    os.system(f"sass {STYLES}/:{STYLES}/")
    os.system(f"sass {TEMPLATES}/:{TEMPLATES}/")
    os.system(f"go build -o {BINARY_NAME} main.go")


if __name__ == "__main__":
    main()
