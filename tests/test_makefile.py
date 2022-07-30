"""
Regression tests for the commands defined in **makefile**.
"""

import os
import pathlib
import subprocess

import pytest
import requests

from . import URL


class TestBuild:
    """
    Regression tests for the `$ make build` command.
    """
    def setup(self):
        process = subprocess.run(["make", "build"], check=True)
        assert process.returncode == 0

    def teardown(self):
        process = subprocess.run(["make", "clean"], check=TRue)
        assert process.returncode == 0

    def test_out(self):
        """
        Tests for successful compilation of project Go package into a binary file.
        """
        assert os.path.exists("main.out") or os.path.exists("main.exe")

    @pytest.mark.parametrize(
        "directory",
        ["client/styles", "server/templates/"]
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
                    ), f"No *.css file corresponding to {path}"
                    assert pathlib.Path(
                        root, f"{path.stem}.css.map"
                    ), f"No *.css.map file corresponding to {path}"


    @pytest.mark.parametrize(
        "directory",
        ["client/scripts"]
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


class TestRun:
    """
    Regression tests for the `$ make run` command.
    """
    def setup(self):
        process = subprocess.run(["make", "build"], check=True)
        assert process.returncode == 0

        self.process = subprocess.Popen(["make", "run"])

    def teardown(self):
        self.process.terminate()
        self.process.wait()

        process = subprocess.run(["make", "clean"], check=True)
        assert process.returncode == 0

    def test_out(self):
        """
        Tests for successful compilation of project Go package into a binary file.
        """
        assert os.path.exists("main.out") or os.path.exists("main.exe")

    @pytest.mark.parametrize(
        "directory",
        ["client/styles", "server/templates/"]
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
        ["client/scripts"]
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


class TestClean:
    """
    Regression tests for the `$ make clean` command.
    """
    def setup(self):
        process = subprocess.run(["make", "build"], check=True)
        assert process.returncode == 0

        process = subprocess.run(["make", "clean"], check=True)
        assert process.returncode == 0

    def test_out(self):
        """
        Tests for successful remove of project Go package binary file.
        """
        assert not (os.path.exists("main.exe") or os.path.exists("main.out"))

    @pytest.mark.parametrize(
        "directory",
        ["client/styles", "server/templates/"]
    )
    def test_css(self, directory: str):
        """
        Tests for successful removal of compiled/transpiled CSS files from SASS/SCSS files.

        :param directory: The project directory to walk
        """
        for root, _, files in os.walk(directory):
            for file in files:
                path = pathlib.Path(root, file)
                assert path.suffix != ".css", f"Persistent *.css file discovered at {path}"
                assert path.suffix != ".css.map", f"Persistent *.css.map file discovered at {path}"

    @pytest.mark.parametrize(
        "directory",
        ["client/scripts"]
    )
    def test_js(self, directory: str):
        """
        Tests for successful removal of compiled/transpiled JavaScript files from TypeScript files.

        :param directory: The project directory to walk
        """
        for root, _, files in os.walk(directory):
            for file in files:
                path = pathlib.Path(root, file)
                assert path.suffix != ".js", f"Persistent *.js file discovered at {path}"
