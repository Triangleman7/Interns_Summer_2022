"""
Regression tests for :py:mod:`make.run`.
"""

import os
import subprocess
import urllib.request

import pytest

from .. import URL
from make import run
from make import constants


class TestRun:
    """
    Regression tests for the `$ python -m make run` command.
    """
    def setup(self):
        self.process = subprocess.Popen(
            ["python" "-m" "make" "run"],
            stdout=subprocess.PIPE,
            stderr=subprocess.PIPE
        )
        self.stdout, self.stderr = process.communicate()

    def teardown(self):
        self.process.terminate()
        assert self.stderr == 0

    def test_out(self):
        """
        Tests for successful compilation of project Go package into a binary file.
        """
        assert constants.BINARY_NAME.exists()

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
        [constants.STYLES, constants.TEMPLATES]
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

    def test_localhost(self):
        """
        """
        with urllib.request.urlopen(URL) as response:
            assert response.code == 200
