var waysToChange = function(n) {
	let res = 0;
	let mod = 1e9 + 7;
	for (let i = 0; i * 25 <= n; i++) {
		let rest = n - i * 25;
		let a = Math.floor(rest / 10);
		let b = Math.floor((rest % 10) / 5);
		res = (res + ((a + 1) * (a + b + 1)) % mod) % mod;
	}
	return res;
};

var n = 5;
console.log(waysToChange(n));
