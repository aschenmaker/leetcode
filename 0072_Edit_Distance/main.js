/*
 * @Author: Aschenmaker 
 * @Date: 2020-04-06 12:07:09 
 * @Last Modified by: Aschenmaker
 * @Last Modified time: 2020-04-06 12:07:47
 */

// 时间复杂度O(mn) 空间复杂度O(mm)

var minDistance = function(word1, word2) {
	// 初始化 DP 矩阵
	let word1Len = word1.length,
		word2Len = word2.length;
	let dp = new Array(word1Len + 1).fill(0).map(() => new Array(word2Len + 1).fill(0));
	// 初始化，行和列
	for (let i = 0; i <= word1Len; i++) {
		dp[i][0] = i;
	}
	for (let i = 0; i <= word2Len; i++) {
		dp[0][i] = i;
	}
	// 计算dp 矩阵
	for (let i = 1; i < dp.length; i++) {
		for (let j = 1; j < dp[i].length; j++) {
			dp[i][j] =
				word1[i - 1] == word2[j - 1]
					? dp[i - 1][j - 1]
					: Math.min(dp[i - 1][j - 1], dp[i][j - 1], dp[i - 1][j]) + 1;
		}
	}
	console.log(dp);
	return dp[word1Len][word2Len];
};

var word1 = 'intention',
	word2 = 'intention';
console.log(minDistance(word1, word2));
