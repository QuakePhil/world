var ctx, conn;

window.onload = function () {
    if (window["WebSocket"]) {
        ctx = document.getElementById("main").getContext("2d")
        conn = new WebSocket("ws://" + document.location.host + "/ws")
        conn.onclose = function (evt) {
            console.log("Connection closed!")
        };
        conn.onmessage = function (evt) {
            if (evt.data) {
                draw(evt.data)
            }
        };
    } else {
        console.log("Your browser does not support WebSockets!")
    }
};

var x, y;

function mousedown(e) {
  x = e.x
  y = e.y
}

function mouseup(e) {
  conn.send([x, y, e.x, e.y].join(" "))
}
