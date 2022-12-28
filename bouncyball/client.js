var spawnX, spawnY;

function mousedown(e) {
  spawnX = e.offsetX;
  spawnY = e.offsetY;
}

function calculatecoordinates(x, y, x2, y2) {
  var xd = x - x2;
  var yd = y - y2;
  return [
    x2,
    y2,
    Math.atan2(yd, xd),
    Math.sqrt(xd * xd + yd * yd),
    document.getElementById("mass").value,
    document.getElementById("charge").value,
  ];
}

function mouseup(e) {
  var yd = e.offsetY - spawnY;
  var xd = e.offsetX - spawnX;
  conn.send(
    calculatecoordinates(e.offsetX, e.offsetY, spawnX, spawnY).join(" ")
  );
  spawnX = undefined;
}

var mouseX, mouseY;

function mousemove(e) {
  mouseX = e.offsetX;
  mouseY = e.offsetY;
  if (mouseY < 20) {
    document.getElementById("settings").style.display = "block";
    window.scrollTo(0, 0);
  }
}

var greeting;

function range(id, min, max) {
  var i = document.createElement("input");
  i.id = id;
  i.type = "range";
  i.min = min;
  i.max = max;
  i.value = (min + max) / 2.0;

  var label = document.createElement("label");
  label.htmlFor = id;
  label.innerText = id; // TODO: ucfirst

  var div = document.createElement("div");
  div.appendChild(i);
  div.appendChild(label);
  return div;
}

function loaded() {
  reset = document.createElement("input");
  reset.type = "reset";

  hide = document.createElement("input");
  hide.type = "button";
  hide.value = "Hide";
  hide.onclick = function () {
    document.getElementById("settings").style.display = "none";
  };

  form = document.createElement("form");
  form.id = "settings";
  form.appendChild(range("mass", 0, 50));
  form.appendChild(range("charge", -1000, 1000));
  form.appendChild(reset);
  form.appendChild(hide);

  document.body.appendChild(form);

  createCanvas(function (canvas) {
    canvas.onmousedown = mousedown;
    canvas.onmouseup = mouseup;
    canvas.onmousemove = mousemove;
  });
  return [window.innerWidth, window.innerHeight].join(" ");
}

function particle(coordinates) {
  // circle
  ctx.beginPath();
  ctx.arc(coordinates[0], coordinates[1], coordinates[4], 0, 2 * Math.PI);
  var r = (coordinates[5] / document.getElementById("charge").max) * 255.0;
  var b = (-coordinates[5] / document.getElementById("charge").max) * 255.0;
  if (coordinates[5] < 0) {
    r = 0;
  } else {
    b = 0;
  }
  ctx.fillStyle = "rgb(" + r + ",0," + b + ")";
  ctx.fill();
  ctx.stroke();
  // line
  ctx.beginPath();
  ctx.moveTo(coordinates[0], coordinates[1]);
  ctx.lineTo(
    parseFloat(coordinates[0]) + Math.cos(coordinates[2]) * coordinates[3],
    parseFloat(coordinates[1]) + Math.sin(coordinates[2]) * coordinates[3]
  );
  ctx.stroke();
}

function coloumb(x, y, x2, y2) {
  // var force = C * charge / distance
  var xd = x - x2;
  var yd = y - y2;
  var r = Math.sqrt(xd * xd + yd * yd);
  var r2 = r * r;
  var force =
    (document.getElementById("charge").value *
      document.getElementById("mass").value) /
    r2;
  console.log("coloumb", x, y, x2, y2, r, force);
  return force;
}

function draw() {
  ctx.clearRect(0, 0, canvas.width, canvas.height);
  frame(function (msg) {
    coordinates = msg.split(" ");
    while (coordinates.length > 0) {
      if (coordinates.length < 6) return;
      particle(coordinates);
      // next
      coordinates.splice(0, 6);
    }
  });

  // circle
  if (spawnX === undefined) {
    ctx.setLineDash([Math.PI, Math.PI]);
    particle(calculatecoordinates(mouseX, mouseY, mouseX, mouseY));
    ctx.setLineDash([]);
  } else {
    particle(calculatecoordinates(mouseX, mouseY, spawnX, spawnY));
  }

  console.log(coloumb(0, 0, mouseX, mouseY));
}
