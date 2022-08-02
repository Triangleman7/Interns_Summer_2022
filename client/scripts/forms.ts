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

        const node = document.createElement("span");
        const datetime = new Date(Date.now());
        if (data["success"]) {
            node.appendChild(document.createTextNode(`${datetime.toString()}: Success.`));
            node.id = "success";
        } else {
            node.appendChild(document.createTextNode(`${datetime.toString()}: Failure.`));
            node.id = "failure";
        }

        let success = document.getElementById("success");
        if (success) {
            formPrimary.removeChild(success);
            formPrimary.reset();
        }
        let failure = document.getElementById("failure");
        if (failure) {
            formPrimary.removeChild(failure);
        }

        formPrimary.appendChild(node);
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
 