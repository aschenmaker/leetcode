const quickSort = function(arr) {
	if (arr.length < 2) {
		return arr;
	}
	var privotIndex = Math.floor(arr.length / 2);
	// 取出 privot 的值，也可以不取出
	var privot = arr.splice(privotIndex, 1)[0];
	var left = [],
		right = [];
	arr.forEach((item) => {
		if (item < privot) {
			left.push(item);
		} else {
			right.push(item);
		}
	});
	return quickSort(left).concat([ privot ], quickSort(right));
};

var arr = [ 0, 1, 2, 6, 7, 5, 9 ];
console.log(quickSort(arr));
