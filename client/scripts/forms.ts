/**
 * Handles events that are triggered when the submit action of a form is invoked.
 */

const formPrimary: HTMLFormElement = document.forms[<any>"primary"];
formPrimary.addEventListener("submit", handleFormPrimary);

/**
 * Handles submission of form element 'form#primary'.
 * Sends to the server a POST request containing the data from each input field.
 * 
 * Disables the default action performed by the browser on form submission.
 * 
 * @param {SubmitEvent} event   Internally passed event received on form submission invocation.
 */
function handleFormPrimary(event: SubmitEvent) {
    let path: string = "forms/primary";
 
    // Disable default action
    event.preventDefault();
 
    // Configure a POST request
    const xhr = new XMLHttpRequest();
    xhr.open("POST", path);
 
    // Prepare form data
    let data = new FormData(formPrimary);
 
    // Send request
    xhr.send(data);

    // Listen for 'load' event
    xhr.onload = () => {
        let data = JSON.parse(xhr.response);
        let datetime = new Date(Date.now());

        //
        let log: string = "";

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
        let tr: HTMLElement;

        //
        tr = document.createElement("tr");
        tr.innerHTML = `<td class="col-timestamp">${datetime.toString()}<span class="expander"></span></td>
        <td class="col-result-status">${data["success"] ? "Success" : "Failure"}</td>`;
        tbody.appendChild(tr);

        //
        tr = document.createElement("tr");
        tr.classList.add("collapsed");
        tr.innerHTML = `<td colspan="2"><div class="log">${log}</div></td>`;
        tbody.appendChild(tr);

        let tbodyFirst = <HTMLElement>document.querySelector("#results-table tbody:nth-of-type(1)")
        table.insertBefore(tbody, tbodyFirst);
    }
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
 