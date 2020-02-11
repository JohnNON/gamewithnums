document.addEventListener("DOMContentLoaded", checkGame);

function checkGame() {
    let csrfToken = document.getElementsByName("Csrf")[0].value;

    if (xmlHttp.readyState == 4 || xmlHttp.readyState == 0) {
        xmlHttp.open("POST", "/private/checkgame", true);
        xmlHttp.setRequestHeader("X-CSRF-Token", csrfToken);
        xmlHttp.onreadystatechange = checkGameServerResponse;
        xmlHttp.send(null);
    }
    else
        setTimeout('check() ', 1000);
}

function checkGameServerResponse() {
    if (xmlHttp.readyState == 4) {
        if (xmlHttp.status == 200) {
            xmlResponse = xmlHttp.responseText;
            try {
                var data = JSON.parse(xmlHttp.responseText);
            } catch (err) {
                console.log(err.message + " in " + xmlHttp.responseText);
                return;
            }

            if (data["val"]) {
                document.getElementById('dialog-rounded').showModal();
            }
        }
        else
            alert("Ошибка доступа к серверу: " + xmlHttp.statusText);
    }
}