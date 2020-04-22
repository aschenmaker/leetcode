/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */

var rightSideView = function(root) {
	if (root == null) return [];
	let arr = [];
	dfs(root, 0, arr);
	return arr;
};
// 使用深度优先搜索
dfs = (root, depth, res) => {
	if (root) {
		if (res.length == depth) {
			res.push(root.val);
		}
		dfs(root.right, depth + 1, res);
		dfs(root.left, depth + 1, res);
	}
};

// 广度优先
bfs = (root, arr) => {
	let res = arr;
	let queue = [];
	queue.push(root);
	while (queue.length != 0) {
		let count = 0;
		let size = queue.length;
		for (let i = 0; i < size; i++) {
			let cur = queue.shift();
			// 从左往右押入下一层的数据，最后的count成为最右边的值
			if (cur.left != null) queue.push(cur.left);
			if (cur.right != null) queue.push(cur.right);
			count = cur.val;
		}

		res.push(count);
	}
};
