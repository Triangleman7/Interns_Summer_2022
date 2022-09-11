/**
 * Handles element animations.
 */

/**
 * Primary handler for popup-style elements.
 * 
 * Popup containers should have the "popup" value in the element class list.
 * 
 * ```
 * <div class="popup">
 *     <span id="..." class="popuptext" >...</span>
 * </div>
 * ```
 * 
 * Popups are considered visible when the value "show" is present in the `div.popup` class list.
 * Popups are considered hiddne when the value "show" is absent in the `div.popup` class list.
 */
class Popup {
    id: string;
    element: HTMLElement

    /**
     * 
     * @param {string} id: The `id` attribute value for the popup descendant element
     */
    constructor(id: string) {
        this.id = id;
        this.element = <HTMLElement>document.getElementById(this.id);
    }

    /**
     * Makes the popup element visible by adding the `show` value to the element class list.
     * 
     * ```
     * <div class="popup show">...</div>
     * ```
     */
    show() {
        this.element.classList.add("show");
    }

    /**
     * Makes the popup element hidden by removing the `show` value from the element class list.
     * 
     * ```
     * <div class="popup">...</div>
     * ```
     */
    hide() {
        this.element.classList.remove("show");
    }
}

const popupFileUpload = new Popup("file-upload-popup");
const popupUploadTimestamp = new Popup("upload-timestamp-popup");
const popupCaptionCasing = new Popup("caption-casing-popup");
const popupCaptionStyling = new Popup("caption-styling-popup");
const popupFormSubmit = new Popup("form-submit-popup");
