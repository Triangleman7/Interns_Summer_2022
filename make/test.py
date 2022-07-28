"""
Imitates the `$ make test` command.
"""

import os


def main():
    # Run Go unit tests
    os.system("go test -v main.go")
