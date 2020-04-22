/*
 * @Author: Aschenmaker 
 * @Date: 2020-04-21 21:41:39 
 * @Last Modified by: Aschenmaker
 * @Last Modified time: 2020-04-21 21:48:41
 */

// 通过排序算法
// var majorityElement = function(nums) {
// 	nums = nums.sort((a, b) => a - b);
// 	let mid = Math.floor(nums.length / 2);
// 	return nums[mid];
// };

// 投票算法
majorityElement = (nums) => {
	let count = 1;
	let majority = nums[0];
	// 从第二个元素开始投票，相同则++,不同则--；
	for (let i = 1; i < nums.length; i++) {
		if (count == 0) {
			majority = nums[i];
		}
		if (nums[i] == majority) {
			count++;
		} else {
			count--;
		}
		console.log(count);
	}
	return majority;
};

var nums = [ 6, 5, 5 ];

console.log(majorityElement(nums));
