function processInput() {
    let textValue: string = document.forms["form"]["form-text-inp"].value;

    // Serialize text field value
    let formData: object = {textInput: textValue};
    // Store serialized data
    localStorage.setItem("formDataJSON", JSON.stringify(formData));

    console.log(localStorage.getItem("formDataJSON"));
}