// 使用深度优先搜索
var levelOrder = function(root) {
	if (!root) return [];
	let res = [];
	dfs(root, 0, res);
	return res;
};

let dfs = (root, deep, res) => {
	if (root) {
		if (!res[deep]) res[deep] = [];
		res[deep].push(root.val);
		dfs(root.left, deep + 1, res);
		dfs(root.right, deep + 1, res);
	}
};

// 使用广度优先搜索

let levelOrder = (root) => {
	if (!root) return [];
	let queue = [ root ];
	let res = [];
	while (queue.length > 0) {
		let len = queue.length;
		let arr = [];
		while (len) {
			let node = queue.shift();
			arr.push(node.val);
			if (node.left) queue.push(node.left);
			if (node.right) queue.push(node.right);
			len--;
		}
		res.push(arr);
	}
	return res;
};
