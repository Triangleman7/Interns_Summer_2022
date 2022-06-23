function processInput() {
    let request = new XMLHttpRequest();
    let data = new FormData()

    let formElement: HTMLFormElement = document.forms[<any>"form"];     // Attribute 'name' value used to get <form> element from DOM (TS expects numerical index; <any> type assertion used)
    let textValue: string = formElement["form-text-inp"].value;

    data.append("text", textValue);

    request.open("POST", "somewhere", true);        // Replace argument 'somwhere' with destination URL
    request.setRequestHeader(
        "Content-type", "applicaiton/x-www-form-urlencoded"
    );
    request.onload = function() {
        console.log(this.responseText);
    }
    request.send(data);
}