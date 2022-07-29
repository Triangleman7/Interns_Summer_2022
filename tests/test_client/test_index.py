"""
Regression tests for **client/index.html**.
"""

import string
import subprocess

from playwright.sync_api import sync_playwright
import pytest
import requests

from .. import URL


HEADLESS = True


@pytest.fixture(scope="module")
def page(request):
    """
    Initializes an automated browser corresponding to the browser type determined by
    `request.param`. Yields a webpage belonging to the appropriate browser.

    :rtype: playwright.sync_api._generated.Page
    """
    with sync_playwright() as play:
        if request.param == "chromium":
            browser = play.chromium.launch(headless=HEADLESS)
        elif request.param == "firefox":
            browser = play.firefox.launch(headless=HEADLESS)
        elif request.param == "webkit":
            browser = play.webkit.launch(headless=HEADLESS)
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

    @pytest.mark.parametrize(
        "query",
        [
            string.ascii_letters,
            string.ascii_lowercase,
            string.ascii_uppercase,
            string.digits,
            string.hexdigits,
            string.octdigits,
            string.punctuation,
            string.printable,
            string.whitespace,
        ]
    )
    def test_form_search(self, page, query: str):
        """
        `header.main form.search`

        :type page: playwright.sync_api._generated.Page
        :param query: The search query to input into the search bar input form
        """
        css = "header.main form.search"
        css_search_input = "div.input-search > input.input-text"

        # Check uniqueness of `css`
        assert len(page.query_selector_all(css)) == 1
        element = page.query_selector(css)

        # Check uniqueness of `css_search_input`
        assert len(page.query_selector_all(css_search_input)) == 1
        search_input = page.query_selector(css_search_input)

        # Submit `query` into search form
        before_url = page.url
        search_input.fill(query)
        search_input.press("Enter")
        page.wait_for_event("submit")
        after_url = page.url
        assert before_url == after_url, query

    def test_nav_main(self, page):
        """
        `header.main nav.main`

        :type page: playwright.sync_api._generated.Page
        """
        css = "header.main nav.main"

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
