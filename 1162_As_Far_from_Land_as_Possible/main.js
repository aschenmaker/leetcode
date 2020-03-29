// var grid = [ [ 1, 0, 1 ], [ 0, 0, 0 ], [ 1, 0, 1 ] ];
// var grid = [ [ 1, 1, 1 ], [ 1, 1, 1 ], [ 1, 1, 1 ] ];
var grid = [
	[ 1, 0, 0, 0, 0, 1, 0, 0, 0, 1 ],
	[ 1, 1, 0, 1, 1, 1, 0, 1, 1, 0 ],
	[ 0, 1, 1, 0, 1, 0, 0, 1, 0, 0 ],
	[ 1, 0, 1, 0, 1, 0, 0, 0, 0, 0 ],
	[ 0, 1, 0, 0, 0, 1, 1, 0, 1, 1 ],
	[ 0, 0, 1, 0, 0, 1, 0, 1, 0, 1 ],
	[ 0, 0, 0, 1, 1, 1, 1, 0, 0, 1 ],
	[ 0, 1, 0, 0, 1, 0, 0, 1, 0, 0 ],
	[ 0, 0, 0, 0, 0, 1, 1, 1, 0, 0 ],
	[ 1, 1, 0, 1, 1, 1, 1, 1, 0, 0 ]
];

var maxDistance = function(grid) {
	let land = [];
	let ocean = [];
	for (let i = 0; i < grid.length; i++) {
		for (let j = 0; j < grid[i].length; j++) {
			if (grid[i][j]) {
				land.push([ i, j ]);
			} else {
				ocean.push([ i, j ]);
			}
		}
	}
	if (land.length == 0 || ocean.length == 0) return -1;
	var max = -1;
	for (let i = 0; i < ocean.length; i++) {
		var min = Infinity;
		for (let j = 0; j < land.length; j++) {
			var distance = countDistance(ocean[i], land[j]);

			if (distance < min) min = distance;
			if (distance == 1) break;
		}

		if (min > max) max = min;
	}
	return max;
};
var countDistance = function(a, b) {
	return Math.abs(a[0] - b[0]) + Math.abs(a[1] - b[1]);
};
console.log(maxDistance(grid));
