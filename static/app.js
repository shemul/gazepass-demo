function Validate(access_token) {
    var xhr = new XMLHttpRequest();
    xhr.withCredentials = true;

    xhr.addEventListener("readystatechange", function () {
        if (this.readyState === 4) {
            console.log(this.responseText);
            if (this.status != 200) {
                document.getElementById("login").hidden = false
            } else {
                document.getElementById("serverInfo").innerText = this.responseText
            }
        }
    });

    xhr.open("GET", "/pvt");
    xhr.setRequestHeader("Authorization", access_token);

    xhr.send();
}

async function gazepassSignIn() {
        var gp = new gazepassjs.default("9866de97-2c85-49ea-9b66-6ddf83d0e6a6");
        var access_token = await gp.getAccessToken();
        Validate(access_token)
}

Validate();