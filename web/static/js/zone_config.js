var zone_id = window.zone_id;
var modify = false;

function display_validation() {
    if (!modify) {
        modify = true;

        // validation button
        button = document.createElement("button");
        button.textContent = "Enregistrer";
        button.className = "bx--btn bx--btn--primary";
        button.type = "submit";
        document.getElementById("validation-div").appendChild(button);

        // cancel button
        link = document.createElement("a");
        link.textContent = "Annuler";
        link.className = "bx--btn bx--btn--secondary";
        link.href = "/parking/zone/" + zone_id;
        document.getElementById("validation-div").appendChild(link);
    }
}

/**
 * These functions are displaying input fields to change values.
 */
function modify_name() {
    display_validation();
    var input = document.createElement("input");
    input.type = "text";
    input.name = "name";
    input.className = "bx--text-input";
    document.getElementById("name-div").appendChild(input);
    document.getElementById("name-button").style.display = "none";
}

function modify_desc() {
    display_validation();
    var input = document.createElement("input");
    input.type = "text";
    input.name = "description";
    input.className = "bx--text-input";
    document.getElementById("desc-div").appendChild(input);
    document.getElementById("desc-button").style.display = "none";
}

function modify_type() {
    display_validation();
    document.getElementById("type-button").style.display = "none";
    var input = document.createElement("select");
    input.className = "bx--select-input";
    input.name = "type";

    var options = {
        "free" : "Gratuit", 
        "paid" : "Payant", 
        "blue" : "Zone bleue"
    }

    for (option in options) {
        var el = document.createElement("option");
        el.value = option;
        el.innerHTML = options[option];
        input.append(el);
    }

    document.getElementById("type-div").appendChild(input);
}

function modify_color() {
    display_validation();
    document.getElementById("color-button").style.display = "none";
    var input = document.createElement("input");
    input.type = "color";
    input.name = "color";
    input.value = window.color;
    document.getElementById("color-div").appendChild(input);
}
