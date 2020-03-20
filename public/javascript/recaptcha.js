function contactUsCallback(page) {
    document.open();
    document.write(page);
    document.close();
}

function recaptchaCallback(token) {
    console.log("received call to recaptchaCallback!\n");
   

    var xhr = new XMLHttpRequest();

    xhr.onreadystatechange = function() {
        if (xhr.readyState === 4) {
          contactUsCallback(xhr.response);
        }
    }

    xhr.open("POST", '/contact_us/recaptcha', true);
    xhr.setRequestHeader('Content-Type', 'application/json');

    xhr.send(JSON.stringify({
        token: token
    }));
    console.log("Made call to backend recaptcha endpoint!\n");
}

