const { ListNode, LinkedList } = require('../ultil/ListNode');

const myReverse = (head, tail) => {
	let pre = null;
	let p = head;
	while (pre !== tail) {
		let next = p.next;
		p.next = pre;
		pre = p;
		p = next;
	}
	return [ tail, head ];
};

const reverseKGroup = (head, k) => {
	let hair = new ListNode(0);
	hair.next = head;
	let pre = hair;

	while (head) {
		let tail = pre;
		// check last elements' length is more than k
		for (let i = 0; i < k; i++) {
			tail = tail.next;
			if (!tail) {
				return hair.next;
			}
		}
		let nex = tail.next;
		[ head, tail ] = myReverse(head, tail);
		pre.next = head;
		tail.next = nex;
		pre = tail;
		head = tail.next;
	}
	return hair.next;
};

let linkedlist = new LinkedList([ 1, 2, 3, 4, 5 ]);
let ret = reverseKGroup(linkedlist.head, 2);
console.log(ret);
while (ret) {
	console.log(ret.val);
	ret = ret.next;
}
