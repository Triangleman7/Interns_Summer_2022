"""
Regression tests for **client/index.html**.
"""

import subprocess

from playwright.sync_api import sync_playwright
import pytest
import requests

from .. import URL


@pytest.fixture(scope="module")
def page(request):
    """
    Initializes an automated browser corresponding to the browser type determined by
    `request.param`. Yields a webpage belonging to the appropriate browser.

    :rtype: playwright.sync_api._generated.Page
    """
    with sync_playwright() as play:
        if request.param == "chromium":
            browser = play.chromium.launch(headless=False)
        elif request.param == "firefox":
            browser = play.firefox.launch(headless=False)
        elif request.param == "webkit":
            browser = play.webkit.launch(headless=False)
        else:
            raise ValueError(f"Could not find matching browser for {request.param}")

        page = browser.new_page()
        page.goto(URL)
        yield page

        page.close()
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
    def setup(self):
        self.process = subprocess.Popen(["make", "run"])

    def teardown(self):
        self.process.terminate()
        self.process.wait()

        process = subprocess.run(["make", "clean"])
        assert process.returncode == 0

    def test_nav_top(self, page):
        """
        `header.main nav.top`

        :type page: playwright.sync_api._generated.Page
        """
        css = "header.main nav.top"

        # Check uniqueness of `css`
        assert len(page.query_selector_all(css)) == 1
        element = page.query_selector(css)

        # Iterate through items in the navigation bar
        for idx, elem in enumerate(element.query_selector_all("ul > li")):
            # Check properties of child `a` element
            assert len(elem.query_selector_all("a")) == 1
            a = elem.query_selector("a")
            assert a.text_content(), idx

            # Check validity of `href` attribute value
            href = a.get_attribute("href")
            assert href is not None
            with requests.get(href) as response:
                assert response.status_code == 200, href
