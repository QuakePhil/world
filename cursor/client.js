function loaded() {
  createCanvas(function (canvas) {
    canvas.onmousemove = function mousemove(e) {
      conn.send([e.offsetX, e.offsetY].join(" "));
    };
  });
}

function draw() {
  ctx.clearRect(0, 0, canvas.width, canvas.height);
  if (latestFrame !== undefined) {
    coordinates = latestFrame.split(" ");
    ctx.beginPath();
    ctx.arc(coordinates[0], coordinates[1], 10, 0, 2 * Math.PI);
    ctx.stroke();
  }
}
