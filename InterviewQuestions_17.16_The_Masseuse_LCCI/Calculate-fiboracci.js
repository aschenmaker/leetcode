/*
 * @Author: Aschenmaker 
 * @Date: 2020-03-24 18:17:52 
 * @Last Modified by: Aschenmaker
 * @Last Modified time: 2020-03-24 19:22:09
 */

// 斐波拉契数列
// 递归方法  0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, …
function recurtionFib(n) {
	if (n < 2) {
		return n;
	} else {
		return recurtionFib(n - 1) + recurtionFib(n - 2);
	}
}

console.log('Answer of recurtion function is  ' + recurtionFib(9)); // 'Answer of recurtion function is  34'

// 动态规划方法
function dynamicFib(n) {
	// 用于存储子问题的值；
	let dp = [ 0, 1, 2 ];
	for (let i = 3; i <= n; i++) {
		dp[i] = dp[i - 1] + dp[i - 2];
	}
	console.log(dp);
	return dp[n - 1];
}
console.log('Answer of dynamic function is  ' + dynamicFib(9)); //'Answer of dynamic function is  1'

// 继续优化
function dynamicFib2(n) {
	if (n <= 0) return 0;
	if (n == 1) return 1;
	let a = 0,
		b = 1;
	for (let i = 2; i <= n; i++) {
		b = a + b;
		a = b - a;
	}
	return b;
}
