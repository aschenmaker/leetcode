var str = '+1';

// var myAtoi = function(str) {
// 	var arr = [ ...str ];
// 	var index = 0;
// 	for (let i = 0; i < arr.length; i++) {
// 		if (arr[i] == ' ') {
// 			continue;
// 		} else {
// 			index = i;
// 			break;
// 		}
// 	}

// 	var res = [];
// 	if (arr[index] == '-') {
// 		while (arr[index + 1] && arr[index + 1].match(/\d/)) {
// 			res.push(arr[index + 1]);
// 			index++;
// 		}
// 		res = res.join('') * -1;
// 	} else if (arr[index] && arr[index].match(/\d/)) {
// 		while (arr[index] && arr[index].match(/\d/)) {
// 			res.push(arr[index]);
// 			index++;
// 		}
// 		res = res.join('') * 1;
// 	} else if (arr[index] == '+') {
// 		while (arr[index + 1] && arr[index + 1].match(/\d/)) {
// 			res.push(arr[index + 1]);
// 			index++;
// 		}
// 		res = res.join('') * 1;
// 	} else {
// 		return 0;
// 	}

// 	if (res <= -2147483648) return -2147483648;
// 	if (res >= 2147483647) return 2147483647;
// 	return res;
// };

myAtoi = (str) => {
	let res = str.trim().match(/^(\-|\+)?\d+/g);
	return res ? Math.max(Math.min(Number(res[0]), 2 ** 31 - 1), -(2 ** 31)) : 0;
};

console.log(myAtoi(str));
