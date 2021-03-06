var seq = '()(())()';

var maxDepthAfterSplit = function(seq) {
	let dep = 0;
	return seq.split('').map((v) => {
		if (v === '(') {
			++dep;
			return dep % 2;
		} else {
			let ans = dep % 2;
			--dep;
			return ans;
		}
	});
};

console.log(maxDepthAfterSplit(seq));
