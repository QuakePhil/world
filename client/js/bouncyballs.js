function draw(obj) {
  coordinates = obj.split(' ');
  // circle
  ctx.beginPath();
  ctx.arc(coordinates[0], coordinates[1], 10, 0, 2 * Math.PI);
  ctx.stroke();
  // line
  ctx.beginPath();
  ctx.moveTo(coordinates[0], coordinates[1]);
  ctx.lineTo(coordinates[2], coordinates[3]);
  ctx.stroke();
}
