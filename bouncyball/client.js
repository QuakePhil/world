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
  ];
}

function mouseup(e) {
  var yd = e.offsetY - spawnY;
  var xd = e.offsetX - spawnX;
  if (xd == 0 && yd == 0) {
    // tap
    document.getElementById("settings").style.display = "block";
    window.scrollTo(0, 0);
  } // drag
  else
    conn.send(
      calculatecoordinates(e.offsetX, e.offsetY, spawnX, spawnY).join(" ")
    );
  spawnX = undefined;
}

var mouseX, mouseY;

function mousemove(e) {
  mouseX = e.offsetX;
  mouseY = e.offsetY;
}

var greeting;

function loaded() {
  massRange = document.createElement("input");
  massRange.id = "mass";
  massRange.type = "range";
  massRange.min = 0;
  massRange.max = 50;
  massRange.value = 50 / 2;

  label = document.createElement("label");
  label.htmlFor = "mass";
  label.innerText = "Mass";

  div = document.createElement("div");
  div.appendChild(massRange);
  div.appendChild(label);

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
  form.appendChild(div);
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

function draw() {
  ctx.clearRect(0, 0, canvas.width, canvas.height);
  if (latestFrame !== undefined) {
    coordinates = latestFrame.split(" ");
    while (coordinates.length > 0) {
      if (coordinates.length < 5) return;
      particle(coordinates);
      // next
      coordinates.splice(0, 5);
    }
  }

  // circle
  if (spawnX === undefined) {
    ctx.setLineDash([Math.PI, Math.PI]);
    particle(calculatecoordinates(mouseX, mouseY, mouseX, mouseY));
    ctx.setLineDash([]);
  } else {
    particle(calculatecoordinates(mouseX, mouseY, spawnX, spawnY));
  }
}
