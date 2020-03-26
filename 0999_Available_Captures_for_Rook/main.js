/*
 * @Author: Aschenmaker 
 * @Date: 2020-03-26 11:06:22 
 * @Last Modified by: Aschenmaker
 * @Last Modified time: 2020-03-26 11:10:45
 */
var board = [
	[ '.', '.', '.', '.', '.', '.', '.', '.' ],
	[ '.', '.', '.', '.', '.', '.', '.', '.' ],
	[ '.', '.', '.', '.', '.', '.', '.', '.' ],
	[ '.', '.', '.', 'R', '.', '.', '.', '.' ],
	[ '.', '.', '.', '.', '.', '.', '.', '.' ],
	[ '.', '.', '.', '.', '.', '.', '.', '.' ],
	[ '.', '.', '.', '.', '.', '.', '.', '.' ],
	[ '.', '.', '.', '.', '.', '.', '.', '.' ]
];

// var numRookCaptures = function(board) {
// 	var row = 0,
// 		column = 0,
// 		counts = 0;
// 	// board.map((rows, index) => {
// 	// 	console.log(rows.indexOf('R'));
// 	// 	if (rows.indexOf('R') != -1) {
// 	// 		console.log(index, rows.indexOf('R'));
// 	// 		row = index;
// 	//         column = rows.indexOf('R');

// 	// 	}
// 	// });
// 	console.log(board[0][1]);
// 	for (let i = 0; i <= 7; i++) {
// 		for (let j = 0; j <= 7; j++) {
// 			if (board[i][j] === 'R') {
// 				row = i;
// 				column = j;
// 				break;
// 			}
// 		}
// 	}
// 	console.log(row, column);
// 	// 向上遍历
// 	var uprow = row,
// 		upcolumn = column;
// 	while (uprow - 1 >= 0 && board[uprow - 1][upcolumn]) {
// 		console.log([ uprow - 1, upcolumn ]);
// 		if (board[uprow - 1][upcolumn] == '.') {
// 			uprow--;
// 		} else if (board[uprow - 1][upcolumn] == 'p') {
// 			counts += 1;
// 			break;
// 		} else if (board[uprow - 1][upcolumn] == 'B') {
// 			break;
// 		}
// 	}
// 	// 向下遍历
// 	var uprow = row,
// 		upcolumn = column;
// 	while (uprow + 1 <= 7 && board[uprow + 1][upcolumn]) {
// 		console.log([ uprow + 1, upcolumn ]);

// 		if (board[uprow + 1][upcolumn] == 'p') {
// 			counts += 1;
// 			break;
// 		}
// 		if (board[uprow + 1][upcolumn] == 'B') {
// 			break;
// 		}
// 		if (board[uprow + 1][upcolumn] == '.') {
// 			uprow++;
// 		}
// 	}
// 	// 向左遍历
// 	var uprow = row,
// 		upcolumn = column;
// 	while (row - 1 >= 0 && board[uprow][upcolumn - 1]) {
// 		console.log([ uprow, upcolumn - 1 ]);

// 		if (board[uprow][upcolumn - 1] == 'p') {
// 			counts += 1;
// 			break;
// 		}
// 		if (board[uprow][upcolumn - 1] == 'B') {
// 			break;
// 		}
// 		if (board[uprow][upcolumn - 1] == '.') {
// 			upcolumn--;
// 		}
// 	}
// 	// 向右遍历
// 	var uprow = row,
// 		upcolumn = column;
// 	while (row + 1 <= 7 && board[uprow][upcolumn + 1]) {
// 		console.log([ uprow, upcolumn + 1 ]);
// 		if (board[uprow][upcolumn + 1] == 'p') {
// 			counts += 1;
// 			break;
// 		}
// 		if (board[uprow][upcolumn + 1] == 'B') {
// 			break;
// 		}
// 		upcolumn++;
// 	}
// 	return counts;
// };

// 使用方向数组
function numRookCaptures(board) {
	let counts = 0,
		row = 0,
		column = 0;
	const dx = [ 0, 1, 0, -1 ];
	const dy = [ 1, 0, -1, 0 ];

	for (let i = 0; i < 8; ++i) {
		for (let j = 0; j < 8; ++j) {
			if (board[i][j] == 'R') {
				row = i;
				column = j;
				break;
			}
		}
	}
	for (let i = 0; i < 4; ++i) {
		for (let step = 0; ; ++step) {
			const tx = row + step * dx[i];
			const ty = column + step * dy[i];
			if (tx < 0 || tx >= 8 || ty < 0 || ty >= 8 || board[tx][ty] == 'B') break;
			if (board[tx][ty] == 'p') {
				counts++;
				break;
			}
		}
	}
	return counts;
}

console.log(numRookCaptures(board));
