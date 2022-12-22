var canvas, ctx, conn, latestFrame;
var fps = 60.0;

window.onload = function() {
    if (window["WebSocket"]) {
        loaded()
        conn = new WebSocket("ws://" + document.location.host + "/ws")
        conn.onmessage = function(e) {
            if (e.data) {
              latestFrame = e.data;
            }
        }

        setInterval(draw, 1000.0/fps);
    } else {
        console.log("Your browser does not support WebSockets!")
    }
};

function createCanvas(cb) {
  canvas = document.createElement("canvas");
  canvas.id = "main";
  canvas.innerText = "Canvas disabled/unsupported"
  cb(canvas)
  document.body.appendChild(canvas);
  ctx = canvas.getContext("2d")
}
