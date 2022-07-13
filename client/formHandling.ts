const formPrimary: HTMLFormElement = document.forms[<any>"primary"];
formPrimary.addEventListener("submit", processFormPrimary);

/**
 * Processes input on submission of `<form name='primary'>` element.
 */
function processFormPrimary(event: Event) {
    // Disable default action
    event.preventDefault();

    // Configure a POST request
    const xhr = new XMLHttpRequest();
    xhr.open("POST", "/");

    // Prepare form data
    let data = new FormData(formPrimary);

    // Set request headers
    xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
    xhr.setRequestHeader("X-Requested-With", "XMLHttpRequest");

    // Send request
    xhr.send(data);

    // Listen for 'load' event
    xhr.onload = () => { console.log(xhr.responseText); }
}