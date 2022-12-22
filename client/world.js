var conn;

window.onload = function() {
    if (window["WebSocket"]) {
        loaded()
        conn = new WebSocket("ws://" + document.location.host + "/ws")
        conn.onmessage = updateFrame
    } else {
        console.log("Your browser does not support WebSockets!")
    }
};
