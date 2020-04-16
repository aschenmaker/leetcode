var merge = function(intervals) {
	if (intervals.length == 0) return [];
	var res = [];
	intervals.sort((a, b) => {
		return a[0] - b[0];
	});
	res.push(intervals[0]);
	for (let i = 0; i < intervals.length; i++) {
		if (intervals[i][0] > res[res.length - 1][1]) {
			res.push(intervals[i]);
		} else if (intervals[i][1] > res[res.length - 1][1]) {
			res[res.length - 1][1] = intervals[i][1];
		}
	}
	return res;
};

var intervals = [ [ 1, 3 ], [ 2, 6 ], [ 8, 10 ], [ 15, 18 ] ];

console.log(merge(intervals));
