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

        // Clear form
        if (data["success"]) {
            formPrimary.reset();
        }

        // Create new node to contain form submission status message
        const node = document.createElement("div");
        node.id = (data["success"] ? "form-success" : "form-failure");
        
        // Format form submission status message
        let status: string = (data["success"] ? "Success" : "Failure");
        let datetime = new Date(Date.now());
        let text: string = `${datetime.toString()}: ${status}`;
        let textnode: Text = document.createTextNode(text);
        
        node.appendChild(textnode)

        // Remove existing #form-success element, if possible
        let success = document.getElementById("form-success");
        if (success) {
            formPrimary.removeChild(success);
        }

        // Remove existing #form-failure element, if possible
        let failure = document.getElementById("form-failure");
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
 