function lcs(word1, word2) {
	if (!word1 || !word2) return '';
	let len1 = word1.length,
		len2 = word2.length,
		longestSubStr = '';
	for (let i = 0; i <= len1; i++) {
		for (let j = 0; j <= len2; j++) {
			var currentStr = '',
				k = 0;
			// 限制在字符串长度之内；两个字符相同，则压入到当前相同的字符串中
			while (i + k < len1 && j + k < len2 && word1[i + k] == word2[j + k]) {
				currentStr += word1[i + k];
				k++;
			}
			if (currentStr.length > longestSubStr.length) {
				longestSubStr = currentStr;
			}
		}
	}
	return longestSubStr;
}

var word1 = 'raven',
	word2 = 'havoc';
// console.log(lcs(word1, word2));

function dpLcs(word1, word2) {
	// 申明一个二维数组
	var dp = [],
		max = 0,
		index = 0;
	for (var i = 0; i < word1.length; i++) {
		dp[i] = [];
		for (var j = 0; j < word2.length; j++) {
			dp[i][j] = 0;
		}
	}
	for (var i = 0; i < word1.length; i++) {
		for (var j = 0; j < word2.length; j++) {
			if (word1[i] == word2[j]) {
				dp[i][j] = dp[i - 1][j - 1] + 1;
			} else {
				dp[i][j] = 0;
			}
			if (max < dp[i][j]) {
				max = dp[i][j];
				index = i;
			}
		}
	}
	console.log(dp, max, index);
	var str = '';
	if (max > 0) {
		for (var i = index - max + 1; str.length < max; i++) {
			str += word1[i];
		}
	}
	return str;
}

console.log(dpLcs(word1, word2));
