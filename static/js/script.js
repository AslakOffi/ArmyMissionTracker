/* function validateForm() {
    var date = document.getElementById("date").value;
    var time = document.getElementById("time").value;
    var location = document.getElementById("location").value;
    var units = document.getElementById("units").value;
    var description = document.getElementById("description").value;
    var email = document.getElementById("email").value;
    var emailError = document.getElementById("emailError");

    if (date === "" || time === "" || location === "" || units === "" || description === "") {
        alert("Veuillez remplir tous les champs du formulaire.");
        return false;
    }

    var emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    if (!emailRegex.test(email)) {
        emailError.style.display = 'inline';
        return false;
    } else {
        emailError.style.display = 'none';
        return true;
    }
}

 */
