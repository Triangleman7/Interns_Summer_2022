"""
Regression tests for **client/index.html**.
"""

import subprocess
import urllib.request

from playwright.sync_api import sync_playwright
import pytest

from . import URL

@pytest.fixture(scope="module")
def page(self, request):
    """
    Initializes an automated browser corresponding to the browser type determined by
    `request.param`. Yields a webpage belonging to the appropriate browser.

    :rtype: playwright.sync_api._generated.Page
    """
    with sync_playwright() as play:
        if request.param == "chromium":
            browser = play.chromium.launch()
        elif request.param == "firefox":
            browser = play.firefox.launch()
        elif request.param == "webkit":
            browser = play.webkit.launch()
        else:
            raise ValueError(f"Could not find matching browser for {request.param}")

        page = browser.new_page()
        yield page

        browser.close()


@pytest.mark.parametrize(
    "page",
    ["chromium", "firefox", "webkit"],
    indirect=True
)
class TestIndex:
    """
    Automated browser testing for `/` (**client/index.html**)
    """
    def setup(self, page):
        """
        :type page: playwright.sync_api._generated.Page
        """
        self.process = subprocess.Popen(["make", "run"], stdout=PIPE, stderr=PIPE)
        self.stdout, self.stderr = process.communicate()

        self.page = page
        page.goto(URL)

    def teardown(self, page):
        """
        :type page: playwright.sync_api._generated.Page
        """
        self.page.close()

        self.process.terminate()
