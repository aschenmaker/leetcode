package main

import (
	"container/list"
	"fmt"
)

//CQueue is a queue achieved by two stacks
type CQueue struct {
	stack1, stack2 *list.List
}

//Constructor return a CQueue instance
func Constructor() CQueue {
	return CQueue{
		stack1: list.New(),
		stack2: list.New(),
	}
}

//AppendTail append a value to tail of the queue
func (c *CQueue) AppendTail(value int) {
	c.stack1.PushBack(value)
}

//DeleteHead delete the head of the queue
func (c *CQueue) DeleteHead() int {
	// if the second stack empty
	if c.stack2.Len() == 0 {
		for c.stack1.Len() > 0 {
			c.stack2.PushBack(c.stack1.Remove(c.stack1.Back()))
		}
	}
	if c.stack2.Len() != 0 {
		e := c.stack2.Back()
		c.stack2.Remove(e)
		return e.Value.(int)
	}
	return -1
}

func main() {
	cq := Constructor()
	cq.AppendTail(1)
	cq.AppendTail(2)
	fmt.Println(cq.DeleteHead())
}
