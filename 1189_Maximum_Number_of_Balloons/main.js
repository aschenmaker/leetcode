var text = 'leetcode';

var maxNumberOfBalloons = function(text) {
	var map = new Map();
	for (i of text) {
		map.set(i, map.has(i) ? map.get(i) + 1 : 1);
	}
	let a = Math.min(map.get('b'), map.get('a'), map.get('n'), map.get('l') / 2, map.get('o') / 2);
	a = parseInt(a);

	if (!a) return 0;
	return a;
};

console.log(maxNumberOfBalloons(text));
