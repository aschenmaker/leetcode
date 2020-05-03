var maxSubArray = function(nums) {
	if (nums.length == 0) return 0;
	let max = nums[0];
	let len = nums.length;
	let before = nums[0];
	for (let i = 1; i < len; i++) {
		let cur = nums[i];
		if (before > 0) {
			cur += before;
		}
		if (cur > max) max = cur;
		before = cur;
	}
	return max;
};

let nums = [ -2, 1 ];
console.log(maxSubArray(nums));
