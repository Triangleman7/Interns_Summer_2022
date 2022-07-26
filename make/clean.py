"""
Imitates the `$ make clean` command.
"""

import os

from .constants import BINARY_NAME, SCRIPTS, STYLES, TEMPLATES


def main():
    os.system(f"go clean")
    for file in SCRIPTS.glob("*.js"):
        os.system(f"rm {file}")
    for file in STYLES.glob("*.css"):
        os.system(f"rm {file}")
    for file in STYLES.glob("*.css.map"):
        os.system(f"rm {file}")
    for file in TEMPLATES.glob("*.css"):
        os.system(f"rm {file}")
    for file in TEMPLATES.glob("*.css.map"):
        os.system(f"rm {file}")
    os.system(f"rm {BINARY_NAME}")


if __name__ == "__main__":
    main()
