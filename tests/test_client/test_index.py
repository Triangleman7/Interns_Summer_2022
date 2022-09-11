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
    default_inputs = {
        "file-upload": pathlib.Path("client", "images", "full-logo.jpg"),
        "upload-timestamp": datetime.datetime.now(),
        "image-scale": 50,
        "image-align": "match-parent",
        "caption-text": __name__,
        "caption-align": "match-parent",
        "caption-casing": "",
        "caption-styling-italic": False,
        "caption-styling-bold": False,
        "caption-styling-underline": False,
        "caption-styling-strikethrough": False,
    }

    def submit_form(self, page, values: dict):
        """
        :type page: playwright.sync_api._generated.Page
        """
        path = pathlib.Path("out", "form-primary")
        output_paths = (
            path / "index.docx",
            path / "index.html",
            path / "styles.scss"
        )

        inputs = {**self.default_inputs, **values}
        elements = {x: page.query_selector(f"form#primary #{x}") for x in self.default_inputs}
        
        elements["file-upload"].set_input_files(inputs["file-upload"])

        elements["upload-timestamp"].click()
        elements["upload-timestamp"].type(inputs["upload-timestamp"].strftime("%m%d%Y"))
        elements["upload-timestamp"].press("Tab")
        elements["upload-timestamp"].type(inputs["upload-timestamp"].strftime("%I%M%p"))

        elements["caption-text"].fill(inputs["caption-text"])

        elements["caption-casing"].select_option(inputs["caption-casing"])

        for x in ("italic", "bold", "underline", "strikethrough"):
            name = f"caption-styling-{x}"
            if inputs[name]:
                elements[name].check()
            else:
                elements[name].uncheck()

        page.click("form#primary #form-submit")

        page.wait_for_event("requestfinished")

        for path in output_paths:
            assert path.exists(), path
            os.remove(path)

    def test_nav_top(self, page):
        """
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
        :type page: playwright.sync_api._generated.Page
        """
        form_fields = [
            "file-upload",
            "upload-timestamp",
            "image-scale",
            "image-align",
            "caption-text",
            "caption-align",
            "caption-casing",
            "caption-styling-italic",
            "caption-styling-bold",
            "caption-styling-underline",
            "caption-styling-strikethrough",
            "form-submit"
        ]

        css = "form#primary"

        # Check uniqueness of `css`
        assert len(page.query_selector_all(css)) == 1
        element = page.query_selector(css)

        # Check <input>/<select> descendant elements
        input_elements = element.query_selector_all("input, select")
        assert len(input_elements) == len(form_fields)
        for idx, elem in enumerate(input_elements):
            name = elem.get_attribute("name")
            assert name is not None, idx
            assert name == form_fields[idx], idx

    @pytest.mark.parametrize(
        "value",
        [
            pathlib.Path("client", "images", "full-logo.jpg")
        ]
    )
    def test_image_upload(self, page, value: str):
        """
        :type page: playwright.sync_api._generated.Page
        :param value:
        """
        css = "form#primary input#file-upload"

        # Check uniqueness of `css`
        assert len(page.query_selector_all(css)) == 1
        element = page.query_selector(css)

        # Check <input> element properties
        assert element.get_attribute("type") == "file"
        assert element.get_attribute("accept") == "image/jpeg"

        self.submit_form(page, {"file-upload": value})

    @pytest.mark.parametrize(
        "value",
        [
            datetime.datetime.now()
        ]
    )
    def test_upload_timestamp(self, page, value: datetime.datetime):
        """
        :type page: playwright.sync_api._generated.Page
        :param value:
        """
        css = "form#primary input#upload-timestamp"

        # Check uniqueness of `css`
        assert len(page.query_selector_all(css)) == 1
        element = page.query_selector(css)

        # Check <input> element properties
        assert element.get_attribute("type") == "datetime-local"

        self.submit_form(page, {"upload-timestamp": value})

    @pytest.mark.parametrize(
        "value",
        range(1, 100, 5)
    )
    def test_image_scale(self, page, value: int):
        """
        :type page: playwright.sync_api._generated.Page
        :param value:
        """
        css = "form#primary input#image-scale"

        # Check uniqueness of `css`
        assert len(page.query_selector_all(css)) == 1
        element = page.query_selector(css)

        # Check <input> element properties
        assert element.get_attribute("type") == "number"

        self.submit_form(page, {"image-scale": value})

    @pytest.mark.parametrize(
        "value",
        [
            "match-parent", "left", "right", "center"
        ]
    )
    def test_image_align(self, page, value: int):
        """
        :type page: playwright.sync_api._generated.Page
        :param value:
        """
        css = "form#primary select#image-align"

        # Check uniqueness of `css`
        assert len(page.query_selector_all(css)) == 1
        element = page.query_selector(css)

        # Check <select> element <option> descendant elements
        option_elements = element.query_selector_all("option")
        assert len(option_elements) == 4
        assert {
            x.get_attribute("value") for x in option_elements
        } == {
            "match-parent", "left", "right", "center"
        }

        self.submit_form(page, {"image-align": value})

    @pytest.mark.parametrize(
        "value",
        [
            string.ascii_letters, string.ascii_lowercase, string.ascii_uppercase,
            string.digits, string.hexdigits, string.octdigits,
            string.punctuation, string.printable, string.whitespace,
        ]
    )
    def test_caption_text(self, page, value: str):
        """
        :type page: playwright.sync_api._generated.Page
        :param value:
        """
        css = "form#primary input#caption-text"

        # Check uniqueness of `css`
        assert len(page.query_selector_all(css)) == 1
        element = page.query_selector(css)

        # Check <input> element properties
        assert element.get_attribute("type") == "text"

        self.submit_form(page, {"caption-text": value})

    @pytest.mark.parametrize(
        "value",
        [
            "match-parent", "start", "end", "left", "right",
            "center", "justify", "justify-all"
        ]
    )
    def test_caption_align(self, page, value: int):
        """
        :type page: playwright.sync_api._generated.Page
        :param value:
        """
        css = "form#primary select#caption-align"

        # Check uniqueness of `css`
        assert len(page.query_selector_all(css)) == 1
        element = page.query_selector(css)

        # Check <select> element <option> descendant elements
        option_elements = element.query_selector_all("option")
        assert len(option_elements) == 8
        assert {
            x.get_attribute("value") for x in option_elements
        } == {
            "match-parent", "start", "end", "left", "right",
            "center", "justify", "justify-all"
        }

        self.submit_form(page, {"caption-align": value})

    @pytest.mark.parametrize(
        "value",
        [
            "", "lower", "upper", "alternating", "camel",
            "dot", "kebab", "opposite", "pascal", "sarcastic",
            "snake", "start", "train"
        ]
    )
    def test_caption_casing(self, page, value: str):
        """
        :type page: playwright.sync_api._generated.Page
        :param value:
        """
        css = "form#primary select#caption-casing"

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

        self.submit_form(page, {"caption-casing": value})

    @pytest.mark.parametrize(
        "value",
        [
            True, False
        ]
    )
    def test_caption_styling_italic(self, page, value: bool):
        """
        :type page: playwright.sync_api._generated.Page
        :param value:
        """
        css = "form#primary input#caption-styling-italic"

        # Check uniqueness of `css`
        assert len(page.query_selector_all(css)) == 1
        element = page.query_selector(css)

        # Check <input> element properties
        assert element.get_attribute("type") == "checkbox"

        self.submit_form(page, {"caption-styling-italic": value})

    @pytest.mark.parametrize(
        "value",
        [
            True, False
        ]
    )
    def test_caption_styling_bold(self, page, value: bool):
        """
        :type page: playwright.sync_api._generated.Page
        :param value:
        """
        css = "form#primary input#caption-styling-bold"

        # Check uniqueness of `css`
        assert len(page.query_selector_all(css)) == 1
        element = page.query_selector(css)

        # Check <input> element properties
        assert element.get_attribute("type") == "checkbox"

        self.submit_form(page, {"caption-styling-bold": value})

    @pytest.mark.parametrize(
        "value",
        [
            True, False
        ]
    )
    def test_caption_styling_underline(self, page, value: bool):
        """
        :type page: playwright.sync_api._generated.Page
        :param value:
        """
        css = "form#primary input#caption-styling-underline"

        # Check uniqueness of `css`
        assert len(page.query_selector_all(css)) == 1
        element = page.query_selector(css)

        # Check <input> element properties
        assert element.get_attribute("type") == "checkbox"

        self.submit_form(page, {"caption-styling-underline": value})

    @pytest.mark.parametrize(
        "value",
        [
            True, False
        ]
    )
    def test_caption_styling_strikethrough(self, page, value: bool):
        """
        :type page: playwright.sync_api._generated.Page
        :param value:
        """
        css = "form#primary input#caption-styling-strikethrough"

        # Check uniqueness of `css`
        assert len(page.query_selector_all(css)) == 1
        element = page.query_selector(css)

        # Check <input> element properties
        assert element.get_attribute("type") == "checkbox"

        self.submit_form(page, {"caption-styling-strikethrough": value})

    def test_form_submit(self, page):
        """
        :type page: playwright.sync_api._generated.Page
        """
        css = "form#primary input#form-submit"

        # Check uniqueness of `css`
        assert len(page.query_selector_all(css)) == 1
        element = page.query_selector(css)

        # Check <input> element properties
        assert element.get_attribute("type") == "submit"

        self.submit_form(page, {})

    def test_info_popups(self, page):
        """
        :type page: playwright.sync_api._generated.Page
        """
        css = "div.popup"
        css_popup_text = "div.popuptext"

        for idx, element in enumerate(page.query_selector_all(css)):
            # Check uniqueness of `css_popup_text` within `element`
            assert len(element.query_selector_all(css_popup_text)) == 1, idx
            elem = element.query_selector(css_popup_text)

            assert not elem.is_visible(), idx       # Popup text is not visible
            element.hover()                         # Hover over `element`
            assert elem.is_visible(), idx           # Popup text is visible
            page.hover("body")                      # Un-hover over `element`
            assert not elem.is_visible(), idx       # Popup text is not visible
