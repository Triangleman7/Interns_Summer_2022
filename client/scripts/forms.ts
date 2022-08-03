/**
 * Handles events that are triggered when the submit action of a form is invoked.
 */


function elementJSONDownload(filename: string, content: string): HTMLElement {
    let a: HTMLElement = document.createElement("a");
    a.classList.add("download-results");
    a.setAttribute(
        "href",
        `data:application/json;charset=utf-8,${encodeURIComponent(content)}`
    );
    a.setAttribute("download", filename);

    let textnode: Text = document.createTextNode("Download");
    a.appendChild(textnode);

    let div: HTMLElement = document.createElement("div");
    div.appendChild(a);

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

function responseFormPrimary(xhr: XMLHttpRequest, form: FormData) {
    let data = JSON.parse(xhr.response);
    let datetime = new Date(Date.now());

    // JSON-seralize form input
    let formObject: any = {};
    form.forEach((value, key) => formObject[key] = value);
    var formJSON = JSON.stringify(formObject);

    //
    let JSONDownload = elementJSONDownload(`form-primary_${Date.now()}.json`, formJSON);

    // Clear form
    if (data["success"]) {
        formPrimary.reset();
    }

    let table = <HTMLElement>document.getElementById("results-table");

    // Create new node to contain form submission information
    let tbody: HTMLElement = document.createElement("tbody");
    tbody.classList.add(
        "results-table-body",
        data["success"] ? "success" : "failure"
    );

    //
    let tr: HTMLElement = document.createElement("tr");

    //
    let tdTimestamp: HTMLElement = document.createElement("td");
    tdTimestamp.classList.add("col-timestamp");
    tdTimestamp.innerText = datetime.toString();
    tr.appendChild(tdTimestamp);

    //
    let tdResultStatus: HTMLElement = document.createElement("td");
    tdResultStatus.classList.add("col-result-status");
    tdResultStatus.innerText = (data["success"] ? "Success" : "Failure");
    tr.appendChild(tdResultStatus);

    //
    let tdFormInput: HTMLElement = document.createElement("td");
    tdFormInput.classList.add("col-form-input");
    tdFormInput.appendChild(JSONDownload);
    tr.appendChild(tdFormInput);

    //
    tbody.appendChild(tr);

    //
    let tbodyFirst = <HTMLElement>document.querySelector("#results-table tbody:nth-of-type(1)")
    table.insertBefore(tbody, tbodyFirst);
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
