var grid = [ [ 2, 2, 2 ], [ 2, 1, 2 ], [ 2, 2, 2 ] ];

var surfaceArea = function(grid) {
	var area = 0;
	for (var i = 0; i <= grid.length - 1; i++) {
		for (var j = 0; j <= grid[i].length - 1; j++) {
			if (grid[i][j] == 0) continue;
			var hide = 0;
			if (i - 1 >= 0 && grid[i - 1][j]) {
				hide += Math.min(grid[i][j], grid[i - 1][j]);
			}
			if (i + 1 <= grid.length - 1 && grid[i + 1][j]) {
				hide += Math.min(grid[i][j], grid[i + 1][j]);
			}
			if (j - 1 >= 0 && grid[i][j - 1]) {
				hide += Math.min(grid[i][j], grid[i][j - 1]);
			}
			if (j + 1 <= grid[i].length - 1 && grid[i][j + 1]) {
				hide += Math.min(grid[i][j], grid[i][j + 1]);
			}
			area += grid[i][j] * 4 + 2 - hide;
		}
	}
	console.log(area);
	return area;
};

console.log(surfaceArea(grid));
