/*
 * @Author: Aschenmaker 
 * @Date: 2020-04-18 10:19:41 
 * @Last Modified by: Aschenmaker
 * @Last Modified time: 2020-04-18 10:35:57
 */

// var maxArea = function(height) {
// 	var s = 0;
// 	for (let i = 0; i < height.length; i++) {
// 		for (let j = 1; j < height.length; j++) {
// 			var curS = (j - i) * Math.min(height[i], height[j]);
// 			if (curS > s) s = curS;
// 		}
// 	}
// 	return s;
// };

// 双指针的方法
var maxArea = function(height) {
	var l = 0,
		r = height.length - 1;
	var res = 0;
	while (l < r) {
		var area = Math.min(height[l], height[r]) * (r - l);
		res = Math.max(res, area);
		if (height[l] <= height[r]) {
			++l;
		} else {
			--r;
		}
	}
	return res;
};

var height = [ 1, 8, 6, 2, 5, 4, 8, 3, 7 ];
console.log(maxArea(height));
