function processInput() {
    let formElement: HTMLFormElement = document.forms[<any>"form"];     // Attribute 'name' value used to get <form> element from DOM (TS expects numerical index; <any> type assertion used)
    let textValue: string = formElement["form-text-inp"].value;

    console.log("form[name='form']", textValue);
}