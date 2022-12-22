var latestFrame;

function updateFrame(frame) {
    if (frame.data) {
      latestFrame = frame.data;
    }
}

var canvas, ctx, conn;
var fps = 60.0;

window.onload = function() {
    if (window["WebSocket"]) {
        loaded()
        conn = new WebSocket("ws://" + document.location.host + "/ws")
        conn.onmessage = updateFrame
        setInterval(draw, 1000.0/fps);
    } else {
        console.log("Your browser does not support WebSockets!")
    }
};
