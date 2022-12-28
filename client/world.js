var canvas, ctx, conn, latestFrame;
var fps = 60.0;

window.onload = () => {
  if (window["WebSocket"]) {
    document.body.onresize = resize;
    greeting = loaded();
    conn = new WebSocket(
      "ws://" + document.location.host + document.location.pathname + "ws"
    );
    conn.onmessage = (e) => {
      if (e.data) latestFrame = e.data;
    };
    if (greeting !== undefined) {
      conn.onopen = () => conn.send(greeting);
    }

    setInterval(draw, 1000.0 / fps);
  } else {
    console.log("Your browser does not support WebSockets!");
  }
};

function frame(cb) {
  if (latestFrame !== undefined) {
    cb(latestFrame);
  }
}

function resize() {
  canvas.width = window.innerWidth;
  canvas.style.width = window.innerWidth;
  canvas.height = window.innerHeight;
  canvas.style.height = window.innerHeight;
}

function createCanvas(cb) {
  canvas = document.createElement("canvas");
  canvas.id = "main";
  canvas.innerText = "Canvas disabled/unsupported";
  resize();
  cb(canvas);
  document.body.appendChild(canvas);
  ctx = canvas.getContext("2d");
}
