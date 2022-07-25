"""
Tests the commands defined in **makefile**.
"""

import os


class TestBuild:
    """
    Tests the `$ make build` command.
    """
    def test(self):
        code: int = os.system("make build")
        assert code == 0

    def teardown(self):
        os.system("make clean")


class TestClean:
    """
    Tests the `$ make clean` command.
    """
    def setup(self):
        os.system("make build")

    def test(self):
        code: int = os.system("make clean")
        assert code == 0
