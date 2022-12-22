function draw(obj) {
    coordinates = obj.split(' ');
    while (coordinates.length > 0) {
        if (coordinates.length < 4) return
        // circle
        ctx.beginPath();
        ctx.arc(coordinates[0], coordinates[1], 10, 0, 2 * Math.PI);
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
        coordinates.splice(0, 4)
    }
}
