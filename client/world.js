var conn;

window.onload = function() {
    if (window["WebSocket"]) {
        conn = new WebSocket("ws://" + document.location.host + "/ws")
        conn.onmessage = loaded()
    } else {
        console.log("Your browser does not support WebSockets!")
    }
};
