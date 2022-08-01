"""
Regression tests for **client/index.html**.
"""

import datetime
import itertools
import os
import pathlib
import string
import subprocess

from playwright.sync_api import sync_playwright
import pytest
import requests

from .. import URL


HEADLESS = False


@pytest.fixture(scope="module")
def page(request):
    """
    Initializes an automated browser corresponding to the browser type determined by
    `request.param`. Yields a webpage belonging to the appropriate browser.

    :rtype: playwright.sync_api._generated.Page
    """
    process = subprocess.run(["make", "build"], check=True)
    assert process.returncode == 0

    with subprocess.Popen(["./main.out"]) as process:
        # Initialize automated browser
        with sync_playwright() as play:
            # Get appropriate web browser type
            if request.param == "chromium":
                browser = play.chromium.launch(headless=HEADLESS)
            elif request.param == "firefox":
                browser = play.firefox.launch(headless=HEADLESS)
            elif request.param == "webkit":
                browser = play.webkit.launch(headless=HEADLESS)
            else:
                raise ValueError(f"Could not find matching browser for {request.param}")

            # Initialize new page and navigate to localhost port
            page = browser.new_page()
            page.goto(URL)
            yield page

            page.close()
            browser.close()

        # Terminate server
        process.terminate()
        process.wait()

    process = subprocess.run(["make", "clean"], check=True)
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
                # TODO: Use once repository is public
                # assert response.status_code == 200, href
                assert response.status_code == 404, href

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
                # TODO: Use once repository is public
                # assert response.status_code == 200, href
                assert response.status_code == 404, href

    def test_form_primary(self, page):
        """
        `form#primary`

        :type page: playwright.sync_api._generated.Page
        """
        css = "form#primary"

        input_fields = [
            "image-upload",
            "image-timestamp",
            "caption-text",
            "caption-casing",
            "italics",
            "bold",
            "underline",
            "strikethrough"
            "submit-form"
        ]

        # Check uniqueness of `css`
        assert len(page.query_selector_all(css)) == 1
        element = page.query_selector(css)

        # Check <input>/<select> descendant elements
        input_elements = element.query_selector_all("input, select")
        assert len(input_elements) == len(input_fields)
        for idx, elem in enumerate(input_elements):
            name = elem.get_attribute("name")
            assert name is not None, idx
            assert name == input_fields[idx], idx

    def test_image_upload(self, page):
        """
        `form#primary input[name='image-upload']`

        :type page: playwright.sync_api._generated.Page
        """
        css = "form#primary input[name='image-upload']"

        # Check uniqueness of `css`
        assert len(page.query_selector_all(css)) == 1
        element = page.query_selector(css)

        # Check <input> element properties
        assert element.get_attribute("type") == "file"
        assert element.get_attribute("accept") == "image/jpeg"

    def test_image_timestamp(self, page):
        """
        `form#primary input[name='image-timestamp']`

        :type page: playwright.sync_api._generated.Page
        """
        css = "form#primary input[name='image-timestamp']"

        # Check uniqueness of `css`
        assert len(page.query_selector_all(css)) == 1
        element = page.query_selector(css)

        # Check <input> element properties
        assert element.get_attribute("type") == "datetime-local"

    def test_caption_text(self, page):
        """
        `form#primary input[name='caption-text']`

        :type page: playwright.sync_api._generated.Page
        """
        css = "form#primary input[name='caption-text']"

        # Check uniqueness of `css`
        assert len(page.query_selector_all(css)) == 1
        element = page.query_selector(css)

        # Check <input> element properties
        assert element.get_attribute("type") == "text"

    def test_caption_casing(self, page):
        """
        `form#primary select[name='caption-casing']`

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

    @pytest.mark.parametrize(
        "name",
        [
            "italics", "bold", "underline", "strikethrough"
        ]
    )
    def test_caption_styling(self, page, name: str):
        """
        `form#primary input[name='(bold|italics|strikethrough|underline)']`

        :type page: playwright.sync_api._generated.Page
        :param name: The `name` attribute of the <input> element
        """
        css = f"form#primary input[name='{name}']"

        # Check uniqueness of items in `css`
        assert len(page.query_selector_all(css)) == 1
        element = page.query_selector(css)

        # Check <input> element properties
        element.get_attribute("type") == "checkbox"

    def test_submit_form(self, page):
        """
        `form#primary input[name='submit-form']`

        :type page: playwright.sync_api._generated.Page
        """
        css = "form#primary input[name='submit-form']"

        # Check uniqueness of `css`
        assert len(page.query_selector_all(css)) == 1
        element = page.query_selector(css)

        # Check <input> element properties
        assert element.get_attribute("type") == "submit"

    def test_info_popups(self, page):
        """
        `div.popup`

        :type page: playwright.sync_api._generated.Page
        :param css: The CSS selector for the info popup element
        """
        css = "div.popup"
        css_popup_text = "span.popuptext"

        for idx, element in enumerate(page.query_selector_all(css)):
            # Check uniqueness of `css_popup_text` within `element`
            assert len(element.query_selector_all(css_popup_text)) == 1, idx
            elem = element.query_selector(css_popup_text)

            assert not elem.is_visible(), idx       # Popup text is not visible
            element.hover()                         # Hover over `element`
            assert elem.is_visible(), idx           # Popup text is visible
            page.hover("body")                      # Un-hover over `element`
            assert not elem.is_visible(), idx       # Popup text is not visible

    @pytest.mark.parametrize(
        "image_upload,image_timestamp,caption_text,caption_casing,caption_styling",
        itertools.product(
            [
                *list(pathlib.Path("client", "images").glob("*.jpg"))
            ],
            [
                datetime.datetime.now()
            ],
            [
                string.ascii_letters, string.ascii_lowercase, string.ascii_uppercase,
                string.digits, string.hexdigits, string.octdigits,
                string.punctuation, string.printable, string.whitespace,
            ],
            [
                "", "lower", "upper", "alternating", "camel",
                "dot", "kebab", "opposite", "pascal", "sarcastic",
                "snake", "start", "train"
            ],
            [
                True, False
            ],
        )
    )
    def test_form_submission(
        self, page,
        image_upload: str, image_timestamp: datetime.datetime,
        caption_text: str, caption_casing: str, caption_styling: bool
    ):
        """
        :type page: playwright.sync_api._generated.Page
        :param image_upload: The value to fill the `input[name='image-upload']` field
        :param image_timestamp: The value to fill the `input[name='image-timestamp']` field
        :param caption_text: The value to fill the `input[name='caption-text']` field
        :param caption_casing: The value to fill the `select[name='caption-casing']` field
        :param caption_styling: The values to fill the `input[name='(bold|italics|strikethrough|underline)']` fields
        """
        path = pathlib.Path("out", "form-primary")
        out_docx = path / "index.docx"
        out_html = path / "index.html"

        input_elements = {
            "image-upload": page.query_selector("form#primary input[name='image-upload']"),
            "image-timestamp": page.query_selector("form#primary input[name='image-timestamp']"),
            "caption-text": page.query_selector("form#primary input[name='caption-text']"),
            "caption-casing": page.query_selector("form#primary select[name='caption-casing']"),
            "italics": page.query_selector("form#primary input[name='italics']"),
            "bold": page.query_selector("form#primary input[name='bold']"),
            "underline": page.query_selector("form#primary input[name='underline']"),
            "strikethrough": page.query_selector("form#primary input[name='strikethrough']"),
            "submit-form": page.query_selector("form#primary input[name='submit-form']")
        }

        input_elements["image-upload"].set_input_files(image_upload)

        input_elements["image-timestamp"].click()
        input_elements["image-timestamp"].type(image_timestamp.strftime("%m%d%Y"))
        input_elements["image-timestamp"].press("Tab")
        input_elements["image-timestamp"].type(image_timestamp.strftime("%I%M%p"))

        input_elements["caption-text"].fill(caption_text)

        input_elements["caption-casing"].select_option(caption_casing)

        for name in ("italics", "bold", "underline", "strikethrough"):
            if caption_styling:
                input_elements[name].check()
            else:
                input_elements[name].uncheck()

        input_elements["submit-form"].click()

        page.wait_for_event("requestfinished")

        assert path.exists()
        assert out_docx.exists()
        assert out_html.exists()
        os.remove(out_docx)
        os.remove(out_html)
