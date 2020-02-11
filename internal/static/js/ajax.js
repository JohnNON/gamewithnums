var xmlHttp = createXmlHttpRequestObject();

function createXmlHttpRequestObject() {
    var xmlHttp;
    if (window.ActiveXObject)
        try {
            xmlHttp = new ActiveXObject("Microsoft.XMLHTTP");
        }
        catch (е) {
            xmlHttp = false;
        }
    else
        try {
            xmlHttp = new XMLHttpRequest();
        }
        catch (е) {
            xmlHttp = false;
        }
    if (!xmlHttp)
        alert("Error creating the XMLHttpRequest object.");
    else
        return xmlHttp;
}

function newgame() {
    let csrfToken = document.getElementsByName("Csrf")[0].value;

    let difficult = encodeURIComponent(document.getElementById("difficult").value);
    let input = document.getElementById("input");
    input.maxLength = difficult;
    input.value = ""

    if (xmlHttp.readyState == 4 || xmlHttp.readyState == 0) {
        xmlHttp.open("POST", "/private/newgame", true);
        xmlHttp.setRequestHeader("X-CSRF-Token", csrfToken);
        xmlHttp.onreadystatechange = newgameServerResponse;
        xmlHttp.send(JSON.stringify({val: difficult}));
    }
    else
        setTimeout('newgame() ', 1000);
}

function newgameServerResponse() {
    if (xmlHttp.readyState == 4) {
        if (xmlHttp.status == 200) {
            xmlResponse = xmlHttp.responseText;
            try {
                var data = JSON.parse(xmlHttp.responseText);
            } catch (err) {
                console.log(err.message + " in " + xmlHttp.responseText);
                return;
            }

            if (data["val"] === "OK") {
                let baseDiv = document.getElementById("rounds");
                baseDiv.innerHTML = '';
                let div = document.createElement('div');
                div.className = "game-round";
                div.innerHTML = "<p>Новая игра началась!</p>";
                baseDiv.append(div);
                baseDiv.scrollTop = baseDiv.scrollHeight;
                let btn = document.getElementById("checkBtn");
                btn.disabled = false;
                btn.className = "nes-btn is-primary";
            } else {
                alert("Какая-то ошибка: " + data["err"]);
            }
        }
        else
            аlеrt("Ошибка доступа к серверу: " + xmlHttp.statusText);
    }
}

function check() {
    let csrfToken = document.getElementsByName("Csrf")[0].value;

    if (xmlHttp.readyState == 4 || xmlHttp.readyState == 0) {
        input = encodeURIComponent(document.getElementById("input").value);
        xmlHttp.open("POST", "/private/check", true);
        xmlHttp.setRequestHeader("X-CSRF-Token", csrfToken);
        xmlHttp.setRequestHeader('Content-Type', 'application/json');
        xmlHttp.onreadystatechange = checkServerResponse;
        xmlHttp.send(JSON.stringify({val: input}));
    }
    else
        setTimeout('check() ', 1000);
}

function checkServerResponse() {
    if (xmlHttp.readyState == 4) {
        if (xmlHttp.status == 200) {
            xmlResponse = xmlHttp.responseText;
            try {
                var data = JSON.parse(xmlHttp.responseText);
            } catch (err) {
                console.log(err.message + " in " + xmlHttp.responseText);
                return;
            }

            let baseDiv = document.getElementById("rounds");
            let div = document.createElement('div');
            div.className = "game-round";
            if (data["status"] === "Win") {
                div.innerHTML = "<p>Вы отгадали правильно!</p> <p>Конец игры.</p>" + "<p>" + data["val"] + "</p>";
                let btn = document.getElementById("checkBtn");
                btn.disabled = true;
                btn.className = "nes-btn is-primary is-disabled";
            } else {
                div.innerHTML = "<p>" + data["val"] + "</p>" + "<p>" + data["code"] + "</p>";
            }
            baseDiv.append(div);
            baseDiv.scrollTop = baseDiv.scrollHeight;
        }
        else
            alert("Ошибка доступа к серверу: " + xmlHttp.statusText);
    }
}

function loadgame() {
    let csrfToken = document.getElementsByName("Csrf")[0].value;

    if (xmlHttp.readyState == 4 || xmlHttp.readyState == 0) {
        xmlHttp.open("POST", "/private/loadgame", true);
        xmlHttp.setRequestHeader("X-CSRF-Token", csrfToken);
        xmlHttp.onreadystatechange = loadGameServerResponse;
        xmlHttp.send(null);
    }
    else
        setTimeout('check() ', 1000);
}

function loadGameServerResponse() {
    if (xmlHttp.readyState == 4) {
        if (xmlHttp.status == 200) {
            xmlResponse = xmlHttp.responseText;
            try {
                var data = JSON.parse(xmlHttp.responseText);
            } catch (err) {
                console.log(err.message + " in " + xmlHttp.responseText);
                return;
            }

            let diff = data[0]["diff"];
            document.getElementById("difficult").value = diff;

            let baseDiv = document.getElementById("rounds");

            let div = document.createElement('div');
            div.className = "game-round";
            div.innerHTML = "<p>Игра загружена!</p>";
            baseDiv.append(div);
            baseDiv.scrollTop = baseDiv.scrollHeight;

            for (var i = 1; i < data.length; i++) {
                let div = document.createElement('div');
                div.className = "game-round";
                div.innerHTML = "<p>" + data[i]["in"] + "</p>" + "<p>" + data[i]["out"] + "</p>";
                baseDiv.append(div);
                baseDiv.scrollTop = baseDiv.scrollHeight;
            }

            let input = document.getElementById("input");
            input.maxLength = diff;
            let btn = document.getElementById("checkBtn");
            btn.disabled = false;
            btn.className = "nes-btn is-primary";
        }
        else
            alert("Ошибка доступа к серверу: " + xmlHttp.statusText);
    }
}

function sound() {
    let sound = document.getElementById("sound");
    if (sound.classList.contains("on")) {
        sound.pause();
        sound.className = "off";
        let img = document.getElementById("soundimg");
        img.alt = "off";
        img.src = "/assets/static/images/sound_off.png";
    } else {
        sound.play();
        sound.className = "on";
        let img = document.getElementById("soundimg");
        img.alt = "on";
        img.src = "/assets/static/images/sound_on.png";
    }
}
