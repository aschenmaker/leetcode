/*
 * @Author: Aschenmaker 
 * @Date: 2020-04-08 11:36:47 
 * @Last Modified by: Aschenmaker
 * @Last Modified time: 2020-04-08 12:31:30
 */

var movingCount = function(m, n, k) {
	const directionArr = [
		[ -1, 0 ], // 上
		[ 0, 1 ], // 右
		[ 1, 0 ], // 下
		[ 0, -1 ] // 左
	];
	// 已经走过的满足条件的坐标点
	let set = new Set([ '0,0' ]);

	// 即将遍历的坐标队列。
	let queue = [ [ 0, 0 ] ];

	// 开始遍历
	while (queue.length) {
		// 结构赋值
		let [ x, y ] = queue.shift();

		// 四个方向进行遍历
		for (let i = 0; i < directionArr.length; i++) {
			let nextX = x + directionArr[i][0];
			let nextY = y + directionArr[i][1];

			// 判断临界值
			if (
				nextX < 0 ||
				nextY < 0 ||
				nextX >= m ||
				nextY >= n ||
				getSum(nextX, nextY) > k ||
				set.has(`${nextX},${nextY}`)
			) {
				continue;
			}

			// 满足的格子加入集合
			set.add(`${nextX},${nextY}`);
			queue.push([ nextX, nextY ]);
		}
	}
	return set.size;
};
// 计算两坐标的数位和
var getSum = function(i, j) {
	let sum = i.toString().split('').concat(j.toString().split(''));
	sum = sum.reduce((res, cur) => {
		res += cur * 1;
		return res;
	}, 0);
	return sum;
};

var m = 16,
	n = 16,
	k = 9;
console.log(movingCount(m, n, k));
