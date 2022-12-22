function loaded() {
    canvas = document.createElement("canvas");
    canvas.id = "main";
    canvas.innerText = "Canvas disabled/unsupported"
    canvas.onmousemove = mousemove
    document.body.appendChild(canvas);
    ctx = canvas.getContext("2d")
}

function mousemove(e) {
    conn.send([e.offsetX, e.offsetY].join(" "))
}

function draw() {
    ctx.clearRect(0, 0, canvas.width, canvas.height);
    if (latestFrame !== undefined) {
        coordinates = latestFrame.split(' ');
        ctx.beginPath();
        ctx.arc(coordinates[0], coordinates[1], 10, 0, 2 * Math.PI);
        ctx.stroke();
    }
}
