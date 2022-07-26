"""
Regression tests for :py:mod:`make.build`.
"""

import os

import pytest

from .. import URL
from make import build
from make import constants


class TestBuild:
    """
    Regression tests for the `$ python -m make build` command.
    """
    def setup(self):
        code = os.system("python -m make build")
        assert code == 0

    def teardown(self):
        code = os.system("python -m make clean")

    def test_out(self):
        """
        Tests for successful compilation of project Go package into a binary file.
        """
        assert os.path.exists(constants.BINARY_NAME)

    @pytest.mark.parametrize(
        "directory",
        [constants.STYLES, constants.TEMPLATES]
    )
    def test_css(self, directory: str):
        """
        Tests for successful compilation/transpilation of SASS/SCSS files into CSS files.

        :param directory: The project directory to walk
        """
        for root, _, files in os.walk(directory):
            for file in files:
                fname, fext = os.path.splitext(file)
                if fext == ".sass" or fext == ".scss":
                    assert os.path.exists(
                        os.path.join(root, f"{fname}.css")
                    ), f"No *.css file corresponding to {os.path.join(root, file)}"
                    assert os.path.exists(
                        os.path.join(root, f"{fname}.css.map")
                    ), f"No *.css.map file corresponding to {os.path.join(root, file)}"

    @pytest.mark.parametrize(
        "directory",
        [constants.SCRIPTS]
    )
    def test_js(self, directory: str):
        """
        Tests for successful compilation/transpilation of TypeScript files into JavaScript files.

        :param directory: The project directory to walk
        """
        for root, _, files in os.walk(directory):
            for file in files:
                fname, fext = os.path.splitext(file)
                if fext == ".ts":
                    assert os.path.exists(
                        os.path.join(root, f"{fname}.js")
                    ), f"No *.js file corresponding to {os.path.join(root, file)}"