var str = 'aabcccccaa';

// function compressor(s) {
// 	var arr = new Array();
// 	var count = 1;
// 	for (let index = 0; index < s.length; index++) {
// 		if (s[index] == s[index + 1]) {
// 			console.log(s[index]);
// 			count++;
// 		} else {
// 			var key = s[index];
// 			var value = count;
// 			arr.push([ key, value ]);
// 			count = 1;
// 		}
// 	}
// 	arr = arr.join('').split(',').join('');
// 	return arr.length < s.length ? arr : s;
// }

var compressString = function(S) {
	let str = S.replace(
		new RegExp(
			'a+|b+|c+|d+|e+|f+|g+|h+|i+|j+|k+|l+|m+|n+|o+|p+|q+|r+|s+|t+|u+|v+|w+|x+|y+|z+|A+|B+|C+|D+|E+|F+|G+|H+|I+|J+|K+|L+|M+|N+|O+|P+|Q+|R+|S+|T+|U+|V+|W+|X+|Y+|Z+',
			'g'
		),
		function(ret) {
			console.log(ret);
			return ret.substring(0, 1) + ret.length;
		}
	);
	return str.length < S.length ? str : S;
};

var s = compressString(str);
console.log(s);
