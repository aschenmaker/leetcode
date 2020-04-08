class ListNode {
	constructor(val) {
		this.val = val;
		this.next = null;
	}
}

class LinkedList {
	constructor(values) {
		this.head = this.createLinkedList(values);
	}
	createLinkedList(value) {
		// console.log('start create');
		let listArr = [];
		for (let i = 0; i < value.length; i++) {
			const element = value[i];
			var node = new ListNode(element);
			listArr.push(node);
		}
		var head = listArr[0];
		for (let i = 0; i < listArr.length; i++) {
			const element = listArr[i];
			if (listArr[i + 1]) {
				element.next = listArr[i + 1];
			}
		}
		return head;
	}
	print() {
		var head = this.head;
		while (head) {
			console.log(head.val);
			head = head.next;
		}
	}
}

module.exports = { ListNode, LinkedList };
