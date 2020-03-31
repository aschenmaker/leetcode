/*
 * @Author: Aschenmaker 
 * @Date: 2020-03-31 11:01:43 
 * @Last Modified by: Aschenmaker
 * @Last Modified time: 2020-03-31 17:28:47
 */

// 直接选择排序 - 时间复杂度 n2
// var sortArray = function(nums) {
// 	for (var i = 0; i < nums.length; i++) {
// 		var min = Infinity;
// 		var minIndex = 0;
// 		for (var j = i; j < nums.length; j++) {
// 			if (nums[j] < min) {
// 				min = nums[j];
// 				minIndex = j;
// 			}
// 		}
// 		// 结构赋值
// 		[ nums[i], nums[minIndex] ] = [ nums[minIndex], nums[i] ];
// 	}
// 	return nums;
// };

// 直接插入排序
var sortArray = function(nums) {
	for (let i = 1; i < nums.length; i++) {
		let temp = nums[i];
		let j = i - 1;
		for (; j >= 0; j--) {
			if (temp >= nums[j]) break;
			nums[j + 1] = nums[j];
		}
		nums[j + 1] = temp;
	}
	return nums;
};

var nums = [ 5, 1, 1, 2, 0, 0 ];
console.log(sortArray(nums));
