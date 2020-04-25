/*
 * @Description: Leetcode daily 0046 return all possible permutations
 * @version: 0.0.1
 * @Company: none
 * @Author: aschenmaker
 * @Date: 2020-04-25 15:33:54
 * @LastEditors: aschenmaker
 * @LastEditTime: 2020-04-25 16:52:16
 * @FilePath: /leetcode/0046_Permutations/permute.js
 */

let permute = (nums) => {
	let len = nums.length;
	let res = [];
	if (len == 0) return res;
	let path = [];
	let used = [];

	dfs(nums, len, 0, path, used, res);
	return res;
};

let dfs = (nums, len, depth, path, used, res) => {
	if (depth == len) {
		// 不能直接使用 path 浅拷贝导致只拷贝了内存地址，结果会随着变化。
		res.push([ ...path ]);
		return;
	}
	for (let i = 0; i < len; i++) {
		if (used[i]) {
			continue;
		}
		path.push(nums[i]);
		used[i] = true;
		dfs(nums, len, depth + 1, path, used, res);
		path.pop();
		used[i] = false;
	}
};

let nums = [ 1, 3 ];
console.log(permute(nums));
