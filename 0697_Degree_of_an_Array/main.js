/*
 * @Author: Aschenmaker 
 * @Date: 2020-03-21 13:33:03 
 * @Last Modified by: Aschenmaker
 * @Last Modified time: 2020-03-21 14:59:25
 */

var nums = [ 1, 2, 1, 3, 3, 4, 3 ];
var findShortestSubArray = function(nums) {
	var map = {};
	for (let i = 0; i < nums.length; i++) {
		map[nums[i]] != undefined ? map[nums[i]].push(i) : (map[nums[i]] = [ i ]);
	}

	if (Object.getOwnPropertyNames(map).length == nums.length) return 1;
	var arr = [],
		max = 0,
		min = Infinity;
	Object.values(map).forEach((value) => {
		if (value.length > 1) arr.push(value);
		max = Math.max(max, value.length);
		arr.push(value);
	});
	arr.forEach((v) => {
		if (v.length == max) {
			min = Math.min(min, v[v.length - 1] - v[0] + 1);
		}
	});
	return min;
};

console.log(findShortestSubArray(nums));
