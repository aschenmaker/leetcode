/*
 * @Author: Aschenmaker 
 * @Date: 2020-03-30 11:31:28 
 * @Last Modified by: Aschenmaker
 * @Last Modified time: 2020-03-30 12:08:23
 */
var n = 10,
	m = 17;
var lastRemaining = function(n, m) {
	var arr = [];
	for (let i = 0; i < n; i++) {
		arr.push(i);
	}
	console.log(arr);
	var cur = 0;
	while (arr.length != 1) {
		cur += m - 1;
		while (cur >= arr.length) cur -= arr.length;
		console.log(cur);
		arr.splice(cur, 1);
		console.log(arr);
	}

	return arr[0];
};
console.log(lastRemaining(n, m));
