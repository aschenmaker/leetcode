var s = 'code';

// 使用map
// var canPermutePalindrome = function(s) {
// 	var map = new Map();
// 	s.split('').map((v) => {
// 		if (map.get(v)) {
// 			map.set(v, map.get(v) + 1);
// 		} else {
// 			map.set(v, 1);
// 		}
// 	});
// 	console.log(map);
// 	var flag = 1;
// 	if (map.length == 1) {
// 		return true;
// 	} else {
// 		map.forEach((value, key) => {
// 			console.log(value, key);
// 			value % 2 == 0 ? '' : flag--;
// 		});
// 		console.log(flag);
// 		if (flag < 0) {
// 			return false;
// 		}
// 	}
// 	return true;
// };

// 使用set
var canPermutePalindrome = function(s) {
	let set = new Set();
	for (var i of s) {
		set.has(i) ? set.delete(i) : set.add(i);
	}
	return set.size <= 1;
};

var a = canPermutePalindrome(s);
console.log(a);
