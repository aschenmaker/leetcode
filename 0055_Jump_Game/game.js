var canJump = function(nums) {
	let distance = 0;
	const len = nums.length;
	for (let i = 0; i < len; i++) {
		if (nums[i] == 0 && distance <= i) break;
		if (i + nums[i] > distance) {
			distance = i + nums[i];
		}
	}
	return distance >= len - 1;
};

var nums = [ 0 ];
console.log(canJump(nums));
