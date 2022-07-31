/**
 * Handles form submissions.
 */


/**
 * Handles submit events sent to a form invoked by form submission.
 */
abstract class Form {
    id: string;
    path: string;
    form: HTMLFormElement;

    /**
     * Searches for the form element with its _id_ attribute set to `id`.
     * Adds to the form element an event listener for **submit** events.
     * Sends POST requests to the URL `path`.
     * 
     * @param id The value of the form element `id` attribute
     * @param path The URL path to which POST requests are to be sent
     */
    constructor(id: string, path: string) {
        this.id = id;
        this.path = path;
        this.form = document.forms[<any>this.id];

        this.form.addEventListener(
            "submit", this.handleRequest
        )
    }

    /**
     * Sends to the server a POST request of the form input data.
     * 
     * The default actions performed by the web browser on form submission are disabled.
     * 
     * @param event Internally passed **submit** event received on form submission.
     */
    handleRequest(event: SubmitEvent): void {
        event.preventDefault();
    }

    /**
     * Handles responses sent by the server after the request is sent.
     */
    handleResponse(): void {}
}


/**
 * Handles form submissions to `form#primary`.
 */
class FormPrimary extends Form {
    xhr: XMLHttpRequest

    /**
     * Input to `form#primary` is sent in a POST request to **\/forms/primary**.
     */
    constructor() {
        super("primary", "/forms/primary");

        this.xhr = new XMLHttpRequest();
    }

    /**
     * Sends a POST request to the server containing the input to `form#primary`.
     * 
     * @param event Internally passed **submit** event received on form submission.
     */
    handleRequest(event: SubmitEvent): void {
        // Disable default actions
        event.preventDefault();

        // Configure a POST request
        this.xhr.open("POST", this.path);

        // Prepare form data
        let data = new FormData(this.form);

        // Send request
        this.xhr.send(data);

        this.xhr.onload = () => this.handleResponse();
    }

    /**
     * Handles responses sent by the server after the request is sent.
     */
    handleResponse(): void {
        console.log(`${this.path}: ${this.xhr.response}`);
    }
}


/**
 * Handles form submissions to `form#search`.
 */
class FormSearch extends Form {

    /**
     * Input to `form#search` is nullified.
     */
    constructor() {
        super("search", "");
    }
}


const formPrimary = new FormPrimary();
const formSearch = new FormSearch();
