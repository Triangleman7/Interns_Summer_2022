"""
Regression tests for :py:mod:`make.clean`.
"""

import os

import pytest

from .. import URL
from make import clean
from make import constants


class TestClean:
    """
    Regression tests for the `$ make clean` command.
    """
    def setup(self):
        os.system("pytHon -m make build")

        code = os.system("python -m make clean")
        assert code == 0

    def test_out(self):
        """
        Tests for successful remove of project Go package binary file.
        """
        assert not constants.BINARY_NAME.exists()

    @pytest.mark.parametrize(
        "directory",
        [constants.STYLES, constants.TEMPLATES]
    )
    def test_css(self, directory: str):
        """
        Tests for successful removal of compiled/transpiled CSS files from SASS/SCSS files.

        :param directory: The project directory to walk
        """
        for root, _, files in os.walk(directory):
            for file in files:
                fname, fext = os.path.splitext(file)
                assert fext != ".css", f"Persistent *.css file discovered at {os.path.join(root, file)}"
                assert fext != ".css.map", f"Persistent *.css.map file discovered at {os.path.join(root, file)}"

    @pytest.mark.parametrize(
        "directory",
        [constants.SCRIPTS]
    )
    def test_js(self, directory: str):
        """
        Tests for successful removal of compiled/transpiled JavaScript files from TypeScript files.

        :param directory: The project directory to walk
        """
        for root, _, files in os.walk(directory):
            for file in files:
                fname, fext = os.path.splitext(file)
                assert fext != ".js", f"Persistent *.js file discovered at {os.path.join(root, file)}"
