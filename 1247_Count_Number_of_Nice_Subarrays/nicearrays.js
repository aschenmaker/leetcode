/*
 * @Author: Aschenmaker 
 * @Date: 2020-04-21 10:16:53 
 * @Last Modified by: Aschenmaker
 * @Last Modified time: 2020-04-21 11:10:27
 */

var numberOfSubarrays = function(nums, k) {
	let n = nums.length,
		counts = new Array(n + 1).fill(0).fill(1, 0, 1),
		oddCount = 0,
		result = 0;
	for (let i = 0; i < n; i++) {
		oddCount += nums[i] & 1;
		result += oddCount >= k ? counts[oddCount - k] : 0;
		counts[oddCount] += 1;
	}
	return result;
};

var nums = [ 1, 1, 2, 1, 1 ],
	k = 2;

console.log(numberOfSubarrays(nums, k));
