function processInput() {
    let formElement: HTMLFormElement = document.forms[<any>"form"];     // Attribute 'name' value used to get <form> element from DOM (TS expects numerical index; <any> type assertion used)
    let textValue: string = formElement["form-text-inp"].value;

    // Serialize text field value
    let formData: object = {textInput: textValue};
    // Locally store serialized data
    localStorage.setItem("formDataJSON", JSON.stringify(formData));

    console.log(localStorage.getItem("formDataJSON"));
}