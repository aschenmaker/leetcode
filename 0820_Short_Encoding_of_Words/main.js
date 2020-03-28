var words = [ 'time', 'me', 'bell', 'time' ];

var minimumLengthEncoding = function(words) {
	let set = new Set(words);

	for (let item of set) {
		for (let i = 1; i < item.length; i++) {
			let target = item.slice(i);

			set.has(target) && set.delete(target);
		}
	}

	let res = 0;
	set.forEach((item) => {
		res += item.length + 1;
	});
	return res;
};

console.log(minimumLengthEncoding(words));
