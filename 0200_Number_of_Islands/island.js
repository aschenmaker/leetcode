/*
 * @Author: Aschenmaker 
 * @Date: 2020-04-20 12:03:23 
 * @Last Modified by: Aschenmaker
 * @Last Modified time: 2020-04-20 12:03:51
 */

// dfs深度优先搜索，将同一岛上的所有元素设置为0；

var numIslands = function(grid) {
	if (grid == null || grid.length == 0) return 0;
	let width = grid[0].length;
	let height = grid.length;
	let num_isLands = 0;
	for (let i = 0; i < height; i++) {
		for (let j = 0; j < width; j++) {
			if (grid[i][j] == 1) {
				++num_isLands;
				dfs(grid, i, j);
			}
		}
	}
	return num_isLands;
};

var dfs = function(grid, i, j) {
	let heigth = grid.length;
	let width = grid[0].length;
	if (i < 0 || j < 0 || i >= heigth || j >= width || grid[i][j] == 0) {
		return;
	}
	grid[i][j] = 0;
	dfs(grid, i, j - 1);
	dfs(grid, i, j + 1);
	dfs(grid, i - 1, j);
	dfs(grid, i + 1, j);
};

console.log(numIslands([]));
