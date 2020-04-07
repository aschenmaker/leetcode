/*
 * @Author: Aschenmaker 
 * @Date: 2020-04-07 10:48:41 
 * @Last Modified by: Aschenmaker
 * @Last Modified time: 2020-04-07 12:07:36
 */

var rotate = function(matrix) {
	var len = matrix.length;
	if (len == 1) return matrix;

	// 按对角线交换
	for (let i = 0; i < len; i++) {
		for (let j = i + 1; j < len; j++) {
			var t = matrix[i][j];
			matrix[i][j] = matrix[j][i];
			matrix[j][i] = t;
		}
	}
	// 对折
	for (let i = 0; i < len; i++) {
		for (let j = 0; j < len / 2; j++) {
			var t = matrix[i][j];
			matrix[i][j] = matrix[i][len - 1 - j];
			matrix[i][len - 1 - j] = t;
		}
	}
	return matrix;
};

var matrix = [ [ 1, 2, 3 ], [ 4, 5, 6 ], [ 7, 8, 9 ] ];

console.log(rotate(matrix));
