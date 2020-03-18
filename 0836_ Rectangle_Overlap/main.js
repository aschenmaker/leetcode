// var rec1 = [ 0, 0, 1, 1 ],
// 	rec2 = [ 1, 0, 2, 1 ];
// var rec1 = [ 0, 0, 2, 2 ],
// 	rec2 = [ 1, 1, 3, 3 ];
var rec1 = [ 7, 8, 13, 15 ],
	rec2 = [ 10, 8, 12, 20 ];

// 解法一 对不相交的情况取反
// var isRectangleOverlap = function(rec1, rec2) {
// 	return !(rec1[2] <= rec2[0] || rec2[2] <= rec1[0] || rec1[3] <= rec2[1] || rec2[3] <= rec1[1]);
// };
// 解法二 抽象相交的条件
var isRectangleOverlap = function(rec1, rec2) {
	return (
		Math.min(rec1[2], rec2[2]) > Math.max(rec1[0], rec2[0]) &&
		Math.min(rec1[3], rec2[3]) > Math.max(rec1[1], rec2[1])
	);
};

var a = isRectangleOverlap(rec1, rec2);
console.log(a);
