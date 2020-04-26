/*
 * @Description: 逐一合并
 * @version: 
 * @Company: none
 * @Author: aschenmaker
 * @Date: 2020-04-26 11:08:00
 * @LastEditors: aschenmaker
 * @LastEditTime: 2020-04-26 16:07:03
 * @FilePath: /leetcode/0023_Merge_k_Sorted_Lists/merge1.js
 */

const { ListNode, LinkedList } = require('../ultil/ListNode.js');

let mergeTwoLists = function(head1, head2) {
	if (!head1 || !head2) return head1 ? head1 : head2;
	let res = new ListNode();
	let head = res;
	while (head1 && head2) {
		if (head1.val <= head2.val) {
			head.next = new ListNode(head1.val);
			head = head.next;
			head1 = head1.next;
		} else {
			head.next = new ListNode(head2.val);
			head = head.next;
			head2 = head2.next;
		}
	}
	head.next = head1 ? head1 : head2;
	return res.next;
};

var mergeKLists = function(lists) {
	if (lists.length == 0) return null;
	let res;
	for (let i = 0; i < lists.length; i++) {
		res = mergeTwoLists(res, lists[i]);
	}
	return res;
};

/**
 * @description: test function 
 * @param {type} 
 * @return: 
 */
let res = mergeKLists(
	[
		// new LinkedList([]).head,
		// new LinkedList([ -1, 5, 11 ]).head,
		// new LinkedList([]).head,
		// new LinkedList([ 6, 10 ]).head
		// new LinkedList([ 1, 4, 5 ]).head,
		// new LinkedList([ 1, 3, 4 ]).head,
		// new LinkedList([ 2, 6 ]).head
	]
);
console.log(res);
while (res.next) {
	console.log(res.val);
	res = res.next;
}
