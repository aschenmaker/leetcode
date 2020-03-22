var A = [ 1, 2, 2 ];

var minIncrementForUnique = function(A) {
	// let set = new Set();
	// var count = 0;
	// for (let i of A) {
	// 	while (set.has(i)) {
	// 		i = i + 1;
	// 		count++;
	// 	}
	// 	set.add(i);
	// 	// console.log(set);
	// }
	// return count;
	A.sort((a, b) => a - b);
	var res = 0;
	var temp = A[0];
	for (var i = 1; i < A.length; i++) {
		if (temp >= A[i]) {
			temp++;
			res += temp - A[i];
		} else {
			temp = A[i];
		}
	}
	return res;
};

console.log(minIncrementForUnique(A));
