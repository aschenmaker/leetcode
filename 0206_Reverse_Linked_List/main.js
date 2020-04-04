const ListNode = require('./ListNode');

var reverseList = function(head) {
	var pre = null;
	var cur = head;
	while (cur != null) {
		var temp = cur.next;
		cur.next = pre;
		pre = cur;
		cur = temp;
	}
	return pre;
};

var reverseList2 = function(head) {
	var pre = null;
	while (head) {
		var next = head.next;
		head.next = pre;
		pre = head;
		head = next;
	}
	return pre;
};

let node1 = new ListNode(1);
let node2 = new ListNode(2);
let node3 = new ListNode(3);
let node4 = new ListNode(4);
let node5 = new ListNode(5);
node1.next = node2;
// node2.next = node3;
// node3.next = node4;
// node4.next = node5;
// node5 = null;

// console.log(reverseList(node1));
console.log(reverseList2(node1));
