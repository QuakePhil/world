var canvas, ctx, conn;

window.onload = function() {
    if (window["WebSocket"]) {
        canvas = document.createElement("canvas");
        canvas.id = "main";
        canvas.innerText = "Canvas disabled/unsupported"
        canvas.onmousedown = mousedown
        canvas.onmouseup = mouseup
        canvas.onmousemove = mousemove
        document.body.appendChild(canvas);
        ctx = canvas.getContext("2d")
        conn = new WebSocket("ws://" + document.location.host + "/ws")
        conn.onmessage = updateFrame
    } else {
        console.log("Your browser does not support WebSockets!")
    }
};
