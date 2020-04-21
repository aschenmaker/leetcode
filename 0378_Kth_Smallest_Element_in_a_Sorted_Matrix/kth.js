/*
 * @Author: Aschenmaker 
 * @Date: 2020-04-20 21:19:18 
 * @Last Modified by: Aschenmaker
 * @Last Modified time: 2020-04-20 21:39:06
 */

// 使用了JavaScript中的转换，显然这么做不是题目的意思。
// var kthSmallest = function(matrix, k) {
// 	var arr = matrix.join(',').split(',').sort((a, b) => a - b);

// 	return arr[k - 1] * 1;
// };

// 这一题主要考察了二分算法。
var kthSmallest = function(matrix, k) {
	let row = matrix.length;
	let col = matrix[0].length;
	let left = matrix[0][0];
	let right = matrix[row - 1][col - 1];
	while (left < right) {
		let mid = Math.floor((left + right) / 2);
		// 在二位矩阵中计算小于mid的个数
		let count = findLessThanMid(matrix, mid, row, col);
		if (count < k) {
			left = mid + 1;
		} else {
			right = mid;
		}
	}
	return right;
};

findLessThanMid = (matrix, mid, row, col) => {
	let i = row - 1;
	let j = 0;
	let count = 0;
	while (i >= 0 && j < col) {
		if (matrix[i][j] <= mid) {
			count += i + 1;
			j++;
		} else {
			i--;
		}
	}
	console.log(count);
	return count;
};

var matrix = [ [ 1, 2 ], [ 1, 3 ] ],
	k = 3;
console.log(kthSmallest(matrix, k));
