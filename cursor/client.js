function loaded() {
  createCanvas(function (canvas) {
    canvas.onmousemove = function mousemove(e) {
      conn.send([e.offsetX, e.offsetY].join(" "));
    };
  });
}

function draw() {
  ctx.clearRect(0, 0, canvas.width, canvas.height);
  frame(function (msg) {
    var coordinates = msg.split(" ");
    var x = parseFloat(coordinates[0]);
    var y = parseFloat(coordinates[1]);
    // re: https://commons.wikimedia.org/wiki/File:Mouse-cursor-hand-pointer.svg
    ctx.beginPath();
    ctx.moveTo(x, y);
    ctx.lineTo(x, y + 18);
    ctx.lineTo(x + 4, y + 14);
    ctx.lineTo(x + 7, y + 20);
    ctx.lineTo(x + 9, y + 19);
    ctx.lineTo(x + 6, y + 13);
    ctx.lineTo(x + 12, y + 13);
    ctx.lineTo(x, y);
    ctx.stroke();
  });
}
