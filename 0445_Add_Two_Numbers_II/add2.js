const { LinkedList, ListNode } = require('../ultil/ListNode');

var addTwoNumbers = function(l1, l2) {
	var l1Stack = [];
	var l2Stack = [];
	while (l1 != null) {
		l1Stack.push(l1.val);
		l1 = l1.next;
	}
	while (l2 != null) {
		l2Stack.push(l2.val);
		l2 = l2.next;
	}
	// console.log(l1Stack, l2Stack);
	let overTen = 0;
	let res = null;
	while (l1Stack.length || l2Stack.length || overTen) {
		var sum = overTen;

		sum += l1Stack.length ? l1Stack.pop() : 0;
		sum += l2Stack.length ? l2Stack.pop() : 0;
		overTen = 0;

		if (sum >= 10) {
			sum = sum % 10;
			overTen = 1;
		}
		var node = new ListNode(sum);
		node.next = res;
		res = node;
	}
	return res;
};

var l1 = new LinkedList([ 7, 2, 4, 3 ]);
var l2 = new LinkedList([ 5, 6, 4 ]);
var print = function(head) {
	while (head) {
		console.log(head.val);
		head = head.next;
	}
};

print(addTwoNumbers(l1.head, l2.head));
