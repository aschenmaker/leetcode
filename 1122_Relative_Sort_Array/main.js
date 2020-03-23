var arr1 = [ 28, 28, 6, 22, 8, 44, 17 ],
	arr2 = [ 22, 28, 8, 6 ];

// var relativeSortArray = function(arr1, arr2) {
// 	let set = new Set(),
// 		arr3 = [];
// 	for (let i of arr1) {
// 		var index = arr2.indexOf(i);
// 		if (index == -1) {
// 			arr3.push(i);
// 		} else {
// 			if (set.has(i)) {
// 				arr2.splice(index, 0, i);
// 			} else {
// 				set.add(i);
// 			}
// 		}
// 	}
// 	arr3.sort((a, b) => a - b);
// 	return [ ...arr2, ...arr3 ];
// };

var relativeSortArray = function(arr1, arr2) {
	return arr1.sort((a, b) => {
		let ia = arr2.indexOf(a);
		let ib = arr2.indexOf(b);
		// 如果 a 和 b 都不在序列中 则 按照 从小到大排列；
		// 即为 a - b > 0 a 于 b 交换位置。
		// console.log(`current a=${a} b=${b} arr:[${arr1}]`);
		if (ia == -1 && ib == -1) {
			return a - b;
		} else if (ia == -1) {
			// 如果 a 不在序列中，b在序列中，则把 b 交换到前面
			return 1;
		} else if (ib == -1) {
			// 如果 a 在序列中，b不在序列中，则保留 a 的位置
			return -1;
		} else {
			// 如果 a 和 b 都在序列中，则返回 序列号小的排在前面
			return ia - ib;
		}
	});
};

console.log(relativeSortArray(arr1, arr2));
