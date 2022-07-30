"""
Regression tests for :py:mod:`make.run`.
"""

import os
import pathlib
import subprocess

import pytest
import requests

from make import constants
from .. import URL


class TestRun:
    """
    Regression tests for the `$ python -m make run` command.
    """
    def setup(self):
        process = subprocess.run(["python", "-m", "make", "build"], check=True)
        assert process.returncode == 0

        self.process = subprocess.Popen(["python", "-m", "make", "run"])

    def teardown(self):
        self.process.terminate()
        self.process.wait()

        process = subprocess.run(["python", "-m", "make", "clean"])
        assert process.returncode == 0

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
                path = pathlib.Path(root, file)
                if path.suffix in (".sass", ".scss"):
                    assert pathlib.Path(
                        root, f"{path.stem}.css"
                    ).exists(), f"No *.css file corresponding to {path}"
                    assert pathlib.Path(
                        root, f"{path.stem}.css.map"
                    ).exists(), f"No *.css.map file corresponding to {path}"

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
                path = pathlib.Path(root, file)
                if path.suffix == ".ts":
                    assert pathlib.Path(
                        root, f"{path.stem}.js"
                    ).exists(), f"No *.js file corresponding to {path}"

    def test_localhost(self):
        """
        """
        with requests.get(URL) as response:
            assert response.status_code == 200
