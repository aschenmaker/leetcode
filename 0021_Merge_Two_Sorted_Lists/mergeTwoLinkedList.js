const { ListNode, LinkedList } = require('../ultil/ListNode.js');

var mergeTwoLists = function(l1, l2) {
	var head1 = l1;
	var head2 = l2;
	var ret = new ListNode();
	var cur = ret;
	while (head1 && head2) {
		if (head1.val <= head2.val) {
			cur.next = new ListNode(head1.val);
			head1 = head1.next;
			cur = cur.next;
		} else {
			cur.next = new ListNode(head2.val);
			head2 = head2.next;
			cur = cur.next;
		}
	}
	if (head1) {
		cur.next = head1;
	} else {
		cur.next = head2;
	}
	return ret.next;
};

var l1 = new LinkedList([ 1, 2, 4 ]);
var l2 = new LinkedList([ 1, 3, 4 ]);
// console.log(l1.head, l2.head);
let ret = mergeTwoLists(l1.head, l2.head);
// console.log(ret);
while (ret) {
	console.log(ret.val);
	ret = ret.next;
}
