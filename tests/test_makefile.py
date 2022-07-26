"""
Regression tests for the commands defined in **makefile**.
"""

import os
import subprocess
import urllib.request

import pytest

from . import URL


class TestBuild:
    """
    Regression tests for the `$ make build` command.
    """
    def setup(self):
        code = os.system("make build")
        assert code == 0

    def teardown(self):
        os.system("make clean")

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
        ["client/styles", "server/templates/"]
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


class TestRun:
    """
    Regression tests for the `$ make run` command.
    """
    def setup(self):
        self.process = subprocess.Popen(
            ["make", "run"], stdout=subprocess.PIPE, stderr=subprocess.PIPE
        )
        self.stdout, self.stderr = process.communicate()

    def teardown(self):
        self.process.terminate()
        assert self.stderr == 0

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
        ["client/styles", "server/templates/"]
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


class TestClean:
    """
    Regression tests for the `$ make clean` command.
    """
    def setup(self):
        os.system("make build")

        code = os.system("make clean")
        assert code == 0

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
                fname, fext = os.path.splitext(file)
                assert fext != ".css", f"Persistent *.css file discovered at {os.path.join(root, file)}"
                assert fext != ".css.map", f"Persistent *.css.map file discovered at {os.path.join(root, file)}"

    @pytest.mark.parametrize(
        "directory",
        ["client/styles", "server/templates/"]
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
