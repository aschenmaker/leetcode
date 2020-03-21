var x = 3,
	y = 5,
	z = 4;

var canMeasureWater = function(x, y, z) {
	if (z == 0) return true;
	if (x == 0 || y == 0) return z == x || z == y;
	var m = Math.max(x, y),
		n = Math.min(x, y);
	if (z > m + n) return false;
	while (m % n != 0) {
		var remain = m % n;
		m = n;
		n = remain;
	}
	return z % n == 0;
};

console.log(canMeasureWater(x, y, z));
