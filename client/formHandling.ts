/**
 * Handles events that are triggered when the submit action of a form is invoked.
 */

// Get element 'form#primary' from HTML DOM
const formPrimary: HTMLFormElement = document.forms[<any>"primary"];
// Listen for form submission events performed on element 'form#primary'
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
    // Disable default action
    event.preventDefault();

    // Configure a POST request
    const xhr = new XMLHttpRequest();
    xhr.open("POST", "/");

    // Prepare form data
    let data = new FormData(formPrimary);

    // Send request
    xhr.send(data);

    // Listen for 'load' event
    xhr.onload = () => { console.log(xhr.responseText); }
}