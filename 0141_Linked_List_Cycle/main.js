const ListNode = require('../ultil/ListNode');

// 快慢指针
var hasCycle = function(head) {
	if (head == null || head.next == null) return false;
	var slow = head,
		fast = head.next;
	while (slow != fast) {
		if (fast == null || fast.next == null) return false;
		slow = slow.next;
		fast = fast.next.next;
	}
	return true;
};

// 生成5个节点的链表
let node1 = new ListNode(1);
let node2 = new ListNode(2);
let node3 = new ListNode(3);
let node4 = new ListNode(4);
let node5 = new ListNode(5);
node1.next = node2;
node2.next = node3;
node3.next = node4;
node4.next = node5;
node5.next = node2;

console.log(hasCycle(node1));
