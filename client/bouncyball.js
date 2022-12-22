var fps = 60.0;
var mass = 10.0;
var spawnX, spawnY;

function mousedown(e) {
    spawnX = e.offsetX
    spawnY = e.offsetY
}

function mouseup(e) {
    var yd = e.offsetY - spawnY
    var xd = e.offsetX - spawnX
    conn.send([spawnX, spawnY, Math.atan2(yd, xd), Math.sqrt(xd*xd + yd*yd), mass].join(" "))
}

var mouseX, mouseY;

function mousemove(e) {
    mouseX = e.offsetX
    mouseY = e.offsetY
}

var latestFrame;

function updateFrame(current) {
    latestFrame = current;
}

function draw() {
    ctx.clearRect(0, 0, canvas.width, canvas.height);
    if (latestFrame !== undefined) {
        coordinates = latestFrame.split(' ');
        while (coordinates.length > 0) {
            if (coordinates.length < 5) return
            // circle
            ctx.beginPath();
            ctx.arc(coordinates[0], coordinates[1], coordinates[4], 0, 2 * Math.PI);
            ctx.stroke();
            // line
            ctx.beginPath();
            ctx.moveTo(coordinates[0], coordinates[1]);
            ctx.lineTo(
                parseFloat(coordinates[0]) + (Math.cos(coordinates[2]) * coordinates[3]),
                parseFloat(coordinates[1]) + (Math.sin(coordinates[2]) * coordinates[3])
            );
            ctx.stroke();
            // next
            coordinates.splice(0, 5)
        }
    }

    // circle
    ctx.beginPath();
    ctx.setLineDash([Math.PI, Math.PI]);
    ctx.arc(mouseX, mouseY, mass, 0, 2 * Math.PI);
    ctx.stroke();
    ctx.setLineDash([]);
}

setInterval(draw, 1000.0/fps);
