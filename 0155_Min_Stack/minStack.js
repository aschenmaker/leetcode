/**
 * initialize your data structure here.
 */
var MinStack = function() {
	this.minStack = [];
	this.minValue = [ Infinity ];
};

/** 
 * @param {number} x
 * @return {void}
 */
MinStack.prototype.push = function(x) {
	this.minStack.push(x);
	this.minValue.push(Math.min(this.minValue[this.minValue.length - 1], x));
};

/**
 * @return {void}
 */
MinStack.prototype.pop = function() {
	this.minStack.pop();
	this.minValue.pop();
};

/**
 * @return {number}
 */
MinStack.prototype.top = function() {
	return this.minStack[this.minStack.length - 1];
};

/**
 * @return {number}
 */
MinStack.prototype.getMin = function() {
	return this.minValue[this.minValue.length - 1];
};

/**
 * Your MinStack object will be instantiated and called as such:
 * var obj = new MinStack()
 * obj.push(x)
 * obj.pop()
 * var param_3 = obj.top()
 * var param_4 = obj.getMin()
 */
