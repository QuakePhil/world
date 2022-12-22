var canvas, ctx, conn;

window.onload = function() {
    if (window["WebSocket"]) {
        canvas = document.getElementById("main")
        ctx = canvas.getContext("2d")
        conn = new WebSocket("ws://" + document.location.host + "/ws")
        conn.onclose = function(evt) {
            console.log("Connection closed!")
        };
        conn.onmessage = function(evt) {
            if (evt.data) {
                updateFrame(evt.data)
            }
        };
    } else {
        console.log("Your browser does not support WebSockets!")
    }
};
