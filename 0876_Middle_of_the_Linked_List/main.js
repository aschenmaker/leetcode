/*
 * @Author: Aschenmaker 
 * @Date: 2020-03-23 10:55:45 
 * @Last Modified by:   Aschenmaker 
 * @Last Modified time: 2020-03-23 10:55:45 
 */

let ListNode = require('./listNode');

// 构建测试用例
var head1 = new ListNode(1);
var head2 = new ListNode(2);
var head3 = new ListNode(3);
var head4 = new ListNode(4);
var head5 = new ListNode(5);
head1.next = head2;
head2.next = head3;
head3.next = head4;
// head4.next = head5;

// 方法一 数组
// 将节点依次压到数组中，计算数组的下标即可。
// var middleNode = function(head) {
// 	let len = 0,
// 		res = [],
// 		newHead = head;
// 	while (newHead) {
// 		res[len++] = newHead;
// 		newHead = newHead.next;
// 	}

// 	return len % 2 == 0 ? res[len / 2] : res[Math.round(len / 2) - 1];
// };

// 方法二 快慢指针法
// 快慢指针法，快指针每次向前两个节点，慢指针每次一个节点。
// 这样快指针到达尾部的时候，慢指针刚好在中间。时间复杂度 O(N), 空间复杂度 O(1);
var middleNode = function(head) {
	let fast = head,
		slow = head;
	while (fast && fast.next) {
		fast = fast.next.next;
		slow = slow.next;
	}
	return slow;
};

console.log(middleNode(head1));
