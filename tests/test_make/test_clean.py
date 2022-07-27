"""
Regression tests for :py:mod:`make.clean`.
"""

import os
import subprocess

import pytest

from make import constants


class TestClean:
    """
    Regression tests for the `$ make clean` command.
    """
    def setup(self):
        process = subprocess.run(["python", "-m", "main", "build"], shell=True)
        assert process.returncode == 0

        process = subprocess.run(["python", "-m", "make", "clean"], shell=True)
        assert process.returncode == 0

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
