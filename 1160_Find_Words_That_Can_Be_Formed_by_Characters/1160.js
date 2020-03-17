// var words = [ 'cat', 'bt', 'hat', 'tree' ];
// var chars = 'atach';
var words = [ 'hello', 'world', 'leetcode' ],
	chars = 'welldonehoneyr';
var countCharacters = function(words, chars) {
	let charsMap = new Map();
	let arr = [];
	chars.split('').map((v) => {
		if (charsMap.has(v)) {
			charsMap.set(v, charsMap.get(v) + 1);
		} else {
			charsMap.set(v, 1);
		}
	});
	console.log(charsMap);

	words.map((v) => {
		let map = new Map();
		v.split('').map((v) => {
			if (map.has(v)) {
				map.set(v, map.get(v) + 1);
			} else {
				map.set(v, 1);
			}
		});

		let flag = 0;
		for (const [ key, value ] of map) {
			if (charsMap.get(key) && charsMap.get(key) >= value) {
				flag = 1;
			} else {
				flag = 0;
				break;
			}
		}
		flag == 1 ? arr.push(v) : (flag = 0);
		map = null;
	});
	console.log(arr);
	return arr.join('').length;
};
var a = countCharacters(words, chars);
console.log(a);
