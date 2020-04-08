/*
 * @Author: Aschenmaker 
 * @Date: 2020-04-06 17:06:37 
 * @Last Modified by: Aschenmaker
 * @Last Modified time: 2020-04-07 10:41:00
 */

const ListNode = require('../ultil/ListNode');

var reverseBetween = function(head, m, n) {
	if (m == n) return head;
	var start = head;
	var index = 1;
	// 找到反转起始节点;
	while (index != m) {
		// 取得反转的前一个节点
		if (index + 1 == m) {
			var reverseStart = head;
		}
		head = head.next;
		index++;
	}
	var reverseEnd = head;
	var pre = null;
	while (head && index != n + 1) {
		var next = head.next;
		head.next = pre;
		pre = head;
		head = next;
		index++;
	}
	if (m == 1) {
		start = pre;
	} else {
		reverseStart.next = pre;
	}
	reverseEnd.next = head;
	// 进行反转
	return start;
};

let node1 = new ListNode(1);
let node2 = new ListNode(2);
let node3 = new ListNode(3);
let node4 = new ListNode(4);
let node5 = new ListNode(5);

node1.next = node2;
node2.next = node3;
node3.next = node4;
node4.next = node5;
node5 = null;

// console.log(reverseBetween(node1, 1, 4));

var showLinkedList = function(head) {
	while (head) {
		console.log(head.val);
		head = head.next;
	}
};

showLinkedList(reverseBetween(node1, 4, 4));
