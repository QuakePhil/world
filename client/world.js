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
                ctx.clearRect(0, 0, canvas.width, canvas.height);
                draw(evt.data)
            }
        };
    } else {
        console.log("Your browser does not support WebSockets!")
    }
};

var originX, originY;

function mousedown(e) {
    originX = e.offsetX
    originY = e.offsetY
}

function mouseup(e) {
    var yd = e.offsetY - originY
    var xd = e.offsetX - originX
    conn.send([originX, originY, Math.atan2(yd, xd), Math.sqrt(xd*xd + yd*yd), 10].join(" "))
}
