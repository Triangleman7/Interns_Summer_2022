"""
Imitates the `$ make test` command.
"""

import os
import pathlib


def main():
    # Run Go unit tests
    os.system(f"go test -v ./{pathlib.Path('resources', 'msword')}")
    os.system(f"go test -v ./{pathlib.Path('server')}")
    os.system(f"go test -v ./{pathlib.Path('server', 'docx')}")
    os.system(f"go test -v ./{pathlib.Path('server', 'html')}")

