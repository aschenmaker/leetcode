/*
 * @Author: Aschenmaker 
 * @Date: 2020-03-24 18:18:37 
 * @Last Modified by: Aschenmaker
 * @Last Modified time: 2020-03-24 22:07:02
 */

var nums = [ 1, 2, 3, 4 ];

var massage = function(nums) {
	if (nums.length == 1) return nums[0];
	if (nums.length == 2) return Math.max(nums[0], nums[1]);
	let dp = [ 0, 0 ];
	for (let i = 2; i < nums.length + 2; i++) {
		dp[i] = Math.max(dp[i - 2] + nums[i - 2], dp[i - 1]);
	}
	return dp[nums.length + 1];
};

console.log(massage(nums));
