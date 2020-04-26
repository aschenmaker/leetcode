/*
 * @Description: 
 * @version: 
 * @Company: none
 * @Author: aschenmaker
 * @Date: 2020-04-26 11:08:00
 * @LastEditors: aschenmaker
 * @LastEditTime: 2020-04-26 12:32:53
 * @FilePath: /leetcode/0023_Merge_k_Sorted_Lists/merge.js
 */

const { ListNode, LinkedList } = require('../ultil/ListNode.js');

// 暴力破解- 非常慢非常差
// var mergeKLists = function(lists) {
// 	let len = lists.length;
// 	let res = new ListNode(null);
// 	if (len == 0) return null;
// 	let curVal = [];
// 	for (let i = 0; i < len; i++) {
// 		if (lists[i]) {
// 			curVal.push(lists[i].val);
// 		} else {
// 			curVal[i] = Infinity;
// 		}
// 	}

// 	let flag = 1;
// 	console.log(curVal);
// 	let head = res;
// 	while (flag) {
// 		let node = new ListNode(Math.min(...curVal));
// 		head.next = node;
// 		head = head.next;
// 		let index = curVal.indexOf(Math.min(...curVal));
// 		if (lists[index]) {
// 			if (lists[index].next == null) {
// 				curVal[index] = Infinity;
// 			} else {
// 				curVal[index] = lists[index].next.val;
// 				lists[index] = lists[index].next;
// 			}
// 		}

// 		if (curVal.every((val) => val == Infinity)) flag = 0;
// 	}
// 	if (res.next.val == Infinity) return null;
// 	return res.next;
// };
let res = mergeKLists([
	new LinkedList([]).head,
	new LinkedList([ -1, 5, 11 ]).head,
	new LinkedList([]).head,
	new LinkedList([ 6, 10 ]).head
	// new LinkedList([ 1, 4, 5 ]).head,
	// new LinkedList([ 1, 3, 4 ]).head,
	// new LinkedList([ 2, 6 ]).head
]);
console.log(res);
while (res.next) {
	console.log(res.val);
	res = res.next;
}
