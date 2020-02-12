function validpass() {
        let inp1 = document.getElementById("password");
        let inp2 = document.getElementById("passwordrepeat");
        if(inp1.value != inp2.value){
            document.getElementById("message").innerHTML = "<p>Введенные пароли не совпадают.</p>";
        }
        document.getElementById("message").innerHTML = "<p>Введенные пароли совпадают.</p>";
}