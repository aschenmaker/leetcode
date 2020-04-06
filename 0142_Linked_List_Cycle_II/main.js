const ListNode = require('../ultil/ListNode');

// 快慢指针
var hasCycle = function(head) {
	if (head == null || head.next == null) return null;
	var slow = head,
		fast = head,
		start = head;

	while (fast != null && fast.next != null) {
		slow = slow.next;
		fast = fast.next.next;
		if (slow == fast) {
			while (start != slow) {
				slow = slow.next;
				start = start.next;
			}
			return slow;
		}
	}

	return null;
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
node5.next = node4;

console.log(hasCycle(node1));
