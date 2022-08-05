/**
 * Handles events that are triggered when the submit action of a form is invoked.
 */

/**
 * 
 * @param filename 
 * @param blob 
 * @returns An anchor element (`<a>`) to download the contents stored in `blob`
 */
function downloadAnchor(filename: string, blob: Blob): HTMLElement {
    let a: HTMLElement = document.createElement("a");
    a.classList.add("download-results");
    a.setAttribute("href", window.URL.createObjectURL(blob));
    a.setAttribute("download", filename);

    let textnode: Text = document.createTextNode("Download");
    a.appendChild(textnode);

    return a;
}

const formPrimary: HTMLFormElement = document.forms[<any>"primary"];
formPrimary.addEventListener("submit", requestFormPrimary);

/**
 * Handles submission of form element 'form#primary'.
 * Sends to the server a POST request containing the data from each input field.
 * 
 * Disables the default action performed by the browser on form submission.
 * 
 * @param {SubmitEvent} event   Internally passed event received on form submission invocation.
 */
function requestFormPrimary(event: SubmitEvent) {
    let path: string = "forms/primary";
 
    // Disable default action
    event.preventDefault();
 
    // Configure a POST request
    const xhr = new XMLHttpRequest();
    xhr.open("POST", path);
 
    // Prepare form data
    let formData = new FormData(formPrimary);
 
    // Send request
    xhr.send(formData);

    // Listen for 'load' event
    xhr.onload = () => { responseFormPrimary(xhr, formData) };
}

/**
 * 
 * @param xhr 
 * @param form 
 */
function responseFormPrimary(xhr: XMLHttpRequest, form: FormData) {
    let data = JSON.parse(xhr.response);
    let datetime = new Date(Date.now());

    // JSON-seralize form input
    let formObject: any = {};
    form.forEach((value, key) => formObject[key] = value);

    // Create blob 
    let blob = new Blob([JSON.stringify(formObject)], {type: "application/json"});
    let anchor: HTMLElement = downloadAnchor(`form-primary_input${Date.now()}.json`, blob);

    // Clear form on successful submission
    if (data["success"]) {
        formPrimary.reset();
    }

    // Create new <tr> element to hold informatino about most recent form submission
    let tr: HTMLElement = document.createElement("tr");
    tr.classList.add("results-table-body", data["success"] ? "success": "failure");
    tr.innerHTML = `<td class="timestamp"></td>
    <td class="result-status"></td>
    <td class="form-input"></td>
    <td class="form-output"></td>`
    tr.querySelector(".timestamp")?.appendChild(document.createTextNode(datetime.toString()));
    tr.querySelector(".result-status")?.appendChild(document.createTextNode(data["success"] ? "Success" : "Failure"));
    tr.querySelector(".form-input")?.appendChild(anchor);

    // Insert new table row as a child of <table> immediately after <thead> element
    let table = <HTMLElement>document.getElementById("results-table");
    let thead = <HTMLElement>document.getElementById("results-table-head");
    table.insertBefore(tr, thead.nextElementSibling);

    requestFormPrimaryZIP()
}

/**
 * 
 */
function requestFormPrimaryZIP() {
    let path: string = "forms/primary/zip";

    // Configure a POST request
    const xhr = new XMLHttpRequest();
    xhr.open("POST", path);

    // Send request
    xhr.send();

    // Listen for 'load' event
    xhr.onload = () => { responseFormPrimaryZIP(xhr) };
}

/**
 * 
 * @param xhr 
 */
function responseFormPrimaryZIP(xhr: XMLHttpRequest) {
    let blob = new Blob([xhr.response], {type: "application/zip"});
    let anchor: HTMLElement = downloadAnchor(`form-primary_output${Date.now()}.zip`, blob);

    let tdFormOutput = <Element>document.querySelector(
        "table#results-table tr:nth-of-type(1) td.form-output"
    );
    tdFormOutput.appendChild(anchor);
}
 
const formSearch: HTMLFormElement = document.forms[<any>"search"];
formSearch.addEventListener("submit", handleFormSearch);
 
/**
 * Handles submission of form element 'form.search'.
 * 
 * Disables the default action performed by the browser on form submission.
 * 
 * @param {SubmitEvent} event   Internally passed event received on form submission invocation.
 */
function handleFormSearch(event: SubmitEvent) {
    // Disable default action
    event.preventDefault();
}
