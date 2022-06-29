/**
 * Processes input on submission of `<form name='primary'>` element.
 */
function processFormPrimary() {
    // Get corresopnding <form> element from DOM
    let element: HTMLFormElement = document.forms[<any>"primary"];     // TS expects numerical index; <any> type assertion used to use index with `name` attribute value

    // Get value of text field
    let textFieldValue: string = element["primary-form-text"].value;

    // Get value of dropdown menu
    let menuValue: string = element["primary-form-text-operation"].value;

    // Debugging
    console.log(`<form name="primary">: Text Field value = "${textFieldValue}"`);
    console.log(`<form name="primary">: Dropdown Menu value = "${menuValue}"`);

    return true;
}