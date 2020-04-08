const { LinkedList, ListNode } = require('../ultil/ListNode');

var addTwoNumbers = function(l1, l2) {
	// console.log(l1, l2);
	let overten = 0;
	let res = new ListNode('head'),
		cur = res;
	while (l1 || l2) {
		let val1 = l1 ? l1.val : 0;
		let val2 = l2 ? l2.val : 0;
		let curVal = val1 + val2 + overten;
		overten = 0;
		if (curVal >= 10) {
			curVal = curVal % 10;
			overten = 1;
		}
		cur.next = new ListNode(curVal);
		cur = cur.next;
		if (l1) l1 = l1.next;
		if (l2) l2 = l2.next;
	}
	if (overten > 0) {
		cur.next = new ListNode(overten);
	}
	return res.next;
};

var l1 = new LinkedList([ 2, 4, 3, 7 ]);
var l2 = new LinkedList([ 5, 6, 4, 5 ]);

var print = function(head) {
	while (head) {
		console.log(head.val);
		head = head.next;
	}
};

print(addTwoNumbers(l1.head, l2.head));
