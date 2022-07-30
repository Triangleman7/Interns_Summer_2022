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
    process = subprocess.Popen(["make", "run"])

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

    process.terminate()
    process.wait()

    process = subprocess.run(["make", "clean"])
    assert process.returncode == 0


@pytest.mark.parametrize(
    "page",
    ["chromium", "firefox", "webkit"],
    indirect=True
)
class TestIndex:
    """
    Automated browser testing for `/` (**client/index.html**)
    """
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
        `header.main form#search`

        :type page: playwright.sync_api._generated.Page
        :param query: The search query to input into the search bar input form
        """
        css = "header.main form#search"
        css_search_input = "div.input-search > input.input-text"

        # Check uniqueness of `css`
        assert len(page.query_selector_all(css)) == 1
        element = page.query_selector(css)

        # Check uniqueness of `css_search_input`
        assert len(element.query_selector_all(css_search_input)) == 1
        search_input = element.query_selector(css_search_input)

        # Submit `query` into search form
        before_url = page.url
        search_input.fill(query)
        search_input.press("Enter")
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

    def test_form_primary(self, page):
        """
        `form#primary`

        :type page: playwright.sync_api._generated.Page
        """
        css = "form#primary"

        input_fields = [
            "image-upload",
            "caption-text",
            "caption-casing",
            "form-submit"
        ]

        # Check uniqueness of `css`
        assert len(page.query_selector_all(css)) == 1
        element = page.query_selector(css)

        # Check <input> descendant elements
        input_elements = element.query_selector_all("input")
        assert len(input_elements) == 4
        for idx, elem in enumerate(input_elements):
            name = elem.get_attribute("name")
            assert name is not None, idx
            assert name == input_fields[idx], idx

    def test_image_upload(self, page):
        """
        `form#primary input[name='image-upload']

        :type page: playwright.sync_api._generated.Page
        """
        css = "form#primary input[name='image-upload']"

        # Check uniqueness of `css`
        assert len(page.query_selector_all(css)) == 1
        element = page.query_selector(css)

        # Check <input> element properties
        assert element.get_attribute("type") == "file"
        assert element.get_attribute("accept") == "image/jpeg"

    def test_image_caption(self, page):
        """
        `form#primary input[name='image-caption']

        :type page: playwright.sync_api._generated.Page
        """
        css = "form#primary input[name='image-caption']"

        # Check uniqueness of `css`
        assert len(page.query_selector_all(css)) == 1
        element = page.query_selector(css)

        # Check <input> element properties
        assert element.get_attribute("type") == "text"

    def test_caption_casing(self, page):
        """
        `form#primary select[name='caption-casing']

        :type page: playwright.sync_api._generated.Page
        """
        css = "form#primary select[name='caption-casing']"

        # Check uniqueness of `css`
        assert len(page.query_selector_all(css)) == 1
        element = page.query_selector(css)

        # Check <select> element <option> descendant elements
        option_elements = element.query_selector_all("option")
        assert len(option_elements) == 13
        assert {
            x.get_attribute("value") for x in option_elements
        } == {
            "", "lower", "upper", "alternating", "camel",
            "dot", "kebab", "opposite", "pascal", "sarcastic",
            "snake", "start", "train"
        }
